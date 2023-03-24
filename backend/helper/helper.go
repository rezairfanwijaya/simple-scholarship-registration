package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetENV(path string) (env map[string]string, err error) {
	env, err = godotenv.Read(path)
	if err != nil {
		return env, fmt.Errorf("failed read env file : %v", err.Error())
	}

	return env, nil
}

func GenerateErrorBinding(err error) (errorBinding []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errorBinding = append(errorBinding, e.Error())
	}

	return errorBinding
}

func HashPassowrd(rawPassword string) (string, error) {
	passByte, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", fmt.Errorf("failed hash password : %v", err.Error())
	}

	return string(passByte), nil
}

func VerifyPassword(rawPassword, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		return fmt.Errorf("failed verify password : %v", err.Error())
	}

	return nil
}

func VerifyLengthPhoneNumber(phoneNumber string) error {
	if len(phoneNumber) < 11 && len(phoneNumber) > 12 {
		return fmt.Errorf("panjang nomor telepon tidak sesuai. Pastikan antara 11 sampai 12 karakter")
	}

	return nil
}
