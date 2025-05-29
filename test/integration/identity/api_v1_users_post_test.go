package identity_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cristiano-pacheco/goflix/test/integration"
)

type PostUsersTestSuite struct {
	suite.Suite
	cmd    *exec.Cmd
	ctx    context.Context
	cancel context.CancelFunc
	client *http.Client
}

func (s *PostUsersTestSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 30*time.Second)

	cmd, err := integration.Bootstrap(s.ctx)
	s.Require().NoError(err)
	s.cmd = cmd

	s.client = &http.Client{Timeout: 10 * time.Second}
}

func (s *PostUsersTestSuite) TearDownTest() {
	if s.cmd != nil {
		integration.Shutdown(s.cmd)
	}
	if s.cancel != nil {
		s.cancel()
	}
}

func TestPostUsersSuite(t *testing.T) {
	suite.Run(t, new(PostUsersTestSuite))
}

func (s *PostUsersTestSuite) TestShouldCreateUserAndReturnStatus201() {
	// Arrange
	requestBody := map[string]string{
		"name":     "Cristiano Pacheco",
		"email":    "chris.spb27@gmail.com",
		"password": "123412341234",
	}

	jsonBody, err := json.Marshal(requestBody)
	s.Require().NoError(err)

	// Act
	req, err := http.NewRequestWithContext(
		s.ctx,
		http.MethodPost,
		"http://localhost:9000/api/v1/users",
		bytes.NewBuffer(jsonBody),
	)
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	// Assert
	s.Equal(http.StatusCreated, resp.StatusCode)
}
