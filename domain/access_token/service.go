package access_token

import "github.com/go-sandbox/bookstore_oauth-api/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	AccessToken, err := s.repository.GetById(accessTokenId)

	if err != nil {
		return nil, err
	}
	return AccessToken, nil
}
