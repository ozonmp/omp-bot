package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	envFile  = ".env"
	tokenEnv = "TOKEN"
	debugEnv = "DEBUG"
)

var (
	ErrTokenNotFound = fmt.Errorf("token not found")
	ErrEmptyToken    = fmt.Errorf("empty token")
)

type Config struct {
	Token string
	Debug bool
}

func loadConfig() (config Config, err error) {
	godotenv.Load(envFile)
	token, err := lookupToken()
	if err != nil {
		return
	}
	debug := lookupDebug()

	config.Token = token
	config.Debug = debug
	return
}

func lookupToken() (token string, err error) {
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

func lookupDebug() (debug bool) {
	debugStr := os.Getenv(debugEnv)
	debug = strings.ToLower(debugStr) == "on"
	return
}
