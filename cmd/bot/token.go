package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	envFile  = ".env"
	tokenEnv = "TOKEN"
)

var (
	ErrTokenNotFound = fmt.Errorf("token not found")
	ErrEmptyToken    = fmt.Errorf("empty token")
)

func lookupToken() (token string, err error) {
	godotenv.Load(envFile)
	token, ok := os.LookupEnv(tokenEnv)
	if !ok {
		err = ErrTokenNotFound
		return
	}
	if token == "" {
		err = ErrEmptyToken
		return
	}
	return
}
