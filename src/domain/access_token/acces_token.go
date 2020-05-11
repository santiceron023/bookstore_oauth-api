package access_token

import (
	"github.com/santiceron023/FUENTE/bookstore_utils-go/rest_errors"
	"strings"
	"time"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"userId"`
	ClientId    int64  `json:"clientId"`
	Expires     int64  `json:"epires"`
}

func (at *AccessToken) Validate() rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires,0).Before(time.Now().UTC())
}

//Client Id ---> which app created the access token
