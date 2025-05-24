package mailer

import (
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/go-mail/mail/v2"
)

func NewDialer(cfg config.Config) *mail.Dialer {
	return mail.NewDialer(cfg.MAIL.Host, cfg.MAIL.Port, cfg.MAIL.Username, cfg.MAIL.Password)
}
