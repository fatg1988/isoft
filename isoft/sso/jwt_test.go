package sso

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func Test_CreateJWT(t *testing.T) {
	tokenString, err := CreateJWT("tom")
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Print("%s\n", tokenString)
	}
}

func Test_ParseJWT(t *testing.T) {
	tokenString, err := CreateJWT("tom")
	if err != nil {
		t.Error(err.Error())
	} else {
		token, errType, err := ParseJWT(tokenString)
		if err != nil {
			t.Error(errType)
		} else {
			for k, v := range token.Header {
				fmt.Printf("%s-%s\n", k, v)
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				var username = claims["username"].(string)
				fmt.Printf("%s\n", username)
			}

		}
	}
}

func Test_ValidateAndParseJWT(t *testing.T) {
	tokenString, err := CreateJWT("tom")
	if err == nil {
		username, _ := ValidateAndParseJWT(tokenString)
		fmt.Printf(username)
	}

}
