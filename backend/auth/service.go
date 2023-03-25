package auth

import (
	"fmt"
	"serkom/helper"

	"github.com/golang-jwt/jwt/v4"
)

type IService interface {
	GenerateToken(userID int) (string, error)
	VerifyToken(token string) (*jwt.Token, error)
}

type Service struct{}

func NewAuthService() *Service {
	return &Service{}
}

func (s *Service) GenerateToken(userID int) (string, error) {
	// claim
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	env, err := helper.GetENV(".env")
	if err != nil {
		return "", err
	}

	// signature
	tokenSigned, err := token.SignedString([]byte(env["SECRET_KEY"]))
	if err != nil {
		return tokenSigned, fmt.Errorf(
			"failed signing token : %v",
			err.Error(),
		)
	}

	return tokenSigned, nil
}

func (s *Service) VerifyToken(token string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signature method")
		}

		env, err := helper.GetENV(".env")
		if err != nil {
			return "", err
		}

		return []byte(env["SECRET_KEY"]), nil
	})

	if err != nil {
		return tokenParsed, fmt.Errorf(
			"failed parse token : %v",
			err.Error(),
		)
	}

	return tokenParsed, nil
}
