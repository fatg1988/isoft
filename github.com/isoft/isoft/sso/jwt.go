package sso

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey = "welcome to wangshubo's blog"
)

// token := jwt.New(jwt.SigningMethodHS256) 和 token.Claims = claims 等同于 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
func CreateJWT(username string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Headers
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"
	// Claims
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["isLogin"] = "isLogin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //1天有效期,过期需要重新登录获取token
	token.Claims = claims
	// Signature
	// 使用自定义字符串加密,并将完整的编码令牌作为字符串
	tokenString, err = token.SignedString([]byte(SecretKey))
	return
}

func ParseJWT(tokenString string) (t *jwt.Token, errType string, err error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, "errInputData", err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, "errExpired", err
			} else {
				// Couldn't handle this token
				return nil, "errInputData", err
			}
		} else {
			// Couldn't handle this token
			return nil, "errInputData", err
		}
	}
	if !token.Valid {
		return nil, "errInputData", err
	}
	return token, "", err
}

func ValidateAndParseJWT(tokenString string) (username string, err error) {
	token, _, err := ParseJWT(tokenString)
	if err == nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			username = claims["username"].(string)
		}
	}
	return
}
