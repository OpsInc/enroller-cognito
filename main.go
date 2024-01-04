package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

type Config struct {
	Cognito     *cognitoidentityprovider.Client
	Username    string
	Password    string
	AppClientID string
	SecretHash  string
}

type GithubConf struct {
	Repo         string
	Organization string
	CognitoToken string
	GithubToken  string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	conf := &Config{
		Cognito:     AWSConnection(),
		Username:    os.Getenv("COGNITO_USER"),
		Password:    os.Getenv("COGNITO_PASSWORD"),
		AppClientID: os.Getenv("COGNITO_CLIENT_ID"),
	}

	tokenID := conf.Signin()

	githubConf := &GithubConf{
		Repo:         "enroller-cognito",
		Organization: "Opsinc",
		CognitoToken: tokenID,
		GithubToken:  os.Getenv("GITHUB_TOKEN"),
	}

	updateStatus := githubConf.UpdateSecrets()
	if updateStatus != int(200) && updateStatus != int(204) {
		log.Fatal("Issue with the UpdateSecrets call, HTTP code received is: ", updateStatus)
	}
}
