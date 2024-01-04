package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (cognitoConfig *Config) AWSConnection() *cognitoidentityprovider.Client {
	// AWS config without authentication since Congito InitiateAuth
	// does not require an AWS account
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("AWS LoadDefaultConfig failed because of error: ", err)
	}

	// Set the Region since we have not authenticated to AWS
	cfg.Region = cognitoConfig.AWSRegion

	return cognitoidentityprovider.NewFromConfig(cfg)
}
