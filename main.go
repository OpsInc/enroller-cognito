package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion   string
	Username    string
	Password    string
	AppClientID string
}

type GithubConf struct {
	Repo         string
	Organization string
	CognitoToken string
	GithubToken  string
}

func main() {
	log.Println("v0.0.2")

	if os.Getenv("GO_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println(err)
		}
	}

	cognitoConf := &Config{
		AWSRegion:   "ca-central-1",
		Username:    os.Getenv("COGNITO_USER"),
		Password:    os.Getenv("COGNITO_PASSWORD"),
		AppClientID: os.Getenv("COGNITO_CLIENT_ID"),
	}

	tokenID := cognitoConf.UserSignin(cognitoConf.AWSConnection())

	// Print to Stdout in order to be fetched by the user or pipeline
	os.Stdout.Write([]byte(tokenID))
}
