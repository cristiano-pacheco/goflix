package service

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/errs"
	"github.com/cristiano-pacheco/goflix/internal/identity/ports"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/logger"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/mailer"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
)

const sendAccountConfirmationEmailTemplate = "account_confirmation.gohtml"
const sendAccountConfirmationEmailSubject = "Account Confirmation"

type SendEmailConfirmationService struct {
	mailerTemplate mailer.Template
	mailer         mailer.SMTPMailer
	userRepository ports.UserRepositoryI
	logger         logger.Logger
	cfg            config.Config
}

var _ ports.SendEmailConfirmationServiceI = (*SendEmailConfirmationService)(nil)

func NewSendEmailConfirmationService(
	mailerTemplate mailer.Template,
	smtpMailer mailer.SMTPMailer,
	userRepository ports.UserRepositoryI,
	logger logger.Logger,
	cfg config.Config,
) *SendEmailConfirmationService {
	return &SendEmailConfirmationService{
		mailerTemplate,
		smtpMailer,
		userRepository,
		logger,
		cfg,
	}
}

func (s *SendEmailConfirmationService) Execute(ctx context.Context, userID uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "SendEmailConfirmationService.Execute")
	defer span.End()

	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		message := "error finding user"
		s.logger.Error(message, "error", err)
		return err
	}

	if user.ConfirmationToken() == nil {
		message := "confirmation token is nil"
		s.logger.Error(message)
		return errs.ErrConfirmationTokenIsNil
	}

	confirmationToken := *user.ConfirmationToken()
	base64Token := base64.StdEncoding.EncodeToString([]byte(confirmationToken))

	// generate the account confirmation link
	accountConfLink := fmt.Sprintf(
		"%s/user/confirmation?id=%d&token=%s",
		s.cfg.App.BaseURL,
		user.ID(),
		base64Token,
	)

	// compile the template
	tplData := struct {
		Name                    string
		AccountConfirmationLink string
	}{
		Name:                    user.Name(),
		AccountConfirmationLink: accountConfLink,
	}

	content, err := s.mailerTemplate.CompileTemplate(sendAccountConfirmationEmailTemplate, tplData)
	if err != nil {
		message := "error compiling template"
		s.logger.Error(message, "error", err)
		return err
	}

	md := mailer.MailData{
		Sender:  s.cfg.MAIL.Sender,
		ToName:  user.Name(),
		ToEmail: user.Email(),
		Subject: sendAccountConfirmationEmailSubject,
		Content: content,
	}

	err = s.mailer.Send(ctx, md)
	if err != nil {
		message := "error sending email"
		s.logger.Error(message, "error", err)
		return err
	}

	return nil
}
