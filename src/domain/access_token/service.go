package access_token

import (
	"github.com/santiceron023/bookstore_oauth-api/src/utils/rest_errors"
	"strings"
)

type Service interface {
	GetById(string) (*AccessToken, *rest_errors.RestError)
}

///no había nada más al crear esto!!!!
type Repository interface {
	GetById(string) (*AccessToken, *rest_errors.RestError)
}

type service struct {
	repository Repository
}

func (s service) GetById(accessTokenId string) (*AccessToken, *rest_errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) ==0 {
		return nil, rest_errors.NewBadRequestError("invalid access token Id")
	}
	accessToken,err := s.repository.GetById(accessTokenId)
	if err != nil{
		return nil,err
	}
	return accessToken, nil
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
