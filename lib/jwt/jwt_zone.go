package jwt

import (
	"gozone/library/conststr"
	"gozone/library/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type ZoneJsonWebTokenHelper struct {}

func (this *ZoneJsonWebTokenHelper) CreateToken(userToken *model.UserToken) (tokenStr string, err error) {

	claims := make(jwt.MapClaims)

	claims["id"] = userToken.Id
	claims["user_name"] = userToken.UserName
	claims["password"] = userToken.Password
	claims["status"] = userToken.Status
	claims["iat"] = time.Now().Unix() - 10
	claims["nbf"] = time.Now().Unix() - 10
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(int64(168))).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenStr, err = token.SignedString([]byte(conststr.AdminXXTEAKey))
	if err != nil {
		return "", err
	}
	return
}

