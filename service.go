package ac

import (
	"net/rpc"

	"github.com/sirupsen/logrus"
)

// Service interface
type Service interface {
	RegisterAndConfirm(name, email, password string, consent bool) (RegisterAndConfirmResponse, error)
	CheckUserAuth(JWT string) (*CheckUserAuthResponse, error)
}

type service struct {
	address string
	logger  *logrus.Logger
}

// New returns a new account service
func New(address string, logger *logrus.Logger) Service {
	return &service{address, logger}
}

func (s *service) RegisterAndConfirm(name, email, password string, consent bool) (RegisterAndConfirmResponse, error) {
	client, err := rpc.DialHTTP("tcp", s.address)
	if err != nil {
		return RegisterAndConfirmResponse{}, err
	}

	var response RegisterAndConfirmResponse
	request := RegisterAndConfirmRequest{
		Name:     name,
		Email:    email,
		Password: password,
		Consent:  consent,
	}
	err = client.Call("RPC.RegisterAndConfirm", request, &response)
	if err != nil {
		return RegisterAndConfirmResponse{}, err
	}

	s.logger.Infof("RPC response %v", response)

	return response, err
}

func (s *service) CheckUserAuth(JWT string) (*CheckUserAuthResponse, error) {
	client, err := rpc.DialHTTP("tcp", s.address)
	if err != nil {
		return &CheckUserAuthResponse{}, err
	}

	var response CheckUserAuthResponse
	request := CheckUserAuthRequest{
		JWT: JWT,
	}
	err = client.Call("RPC.CheckUserAuth", request, &response)
	if err != nil {
		return &CheckUserAuthResponse{}, err
	}

	s.logger.Infof("RPC response %v", response)

	return &response, err
}
