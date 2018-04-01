package main

import (
	"time"

	pb "github.com/craigmpeters/shopper/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("mysupersecretkey123")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(tokenstring string) (*CustomClaims, error)  {

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenstring, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt:expireToken,
			Issuer: "shopper.user",
		},
	}

	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign Token and Return
	return token.SignedString(key)
}
