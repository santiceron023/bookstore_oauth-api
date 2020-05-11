package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAcesTokenConstants(t *testing.T){
	//if expirationTime != 24 {
	//	t.Error("expiration time sohuld be 24")
	//}
	assert.Equal(t,24,expirationTime,"expiration time sohuld be 24")
}
func TestGetNewAccessToken(t *testing.T) {
	accesToken := GetNewAccessToken()
	if accesToken.IsExpired() {
		t.Error("brnad new acces token is nil")
	}

	if accesToken.AccessToken != ""{
		t.Error("new access token cant have id")
	}

	if accesToken.UserId != 0{
		t.Error("new access token cant have user id")
	}
}


func TestAccessTokenIsExpired(t *testing.T) {
 at := AccessToken{}
 if !at.IsExpired(){
 	t.Error("empty access token sohuld be expired")
 }
	at.Expires = time.Now().UTC().Add(3*time.Hour).Unix()
	if at.IsExpired(){
		t.Error("empty access token +3 hor sohuld not be be expired")
	}
}