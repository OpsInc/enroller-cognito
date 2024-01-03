package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func AWSConnection() *cognitoidentityprovider.Client {
	var (
		cfg aws.Config
		err error
	)

	awsProfile := os.Getenv("AWS_PROFILE")

	if awsProfile == "" {
		cfg, err = config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal("Unable to load AWS profile from LAMBDA:", err)
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithSharedConfigProfile(awsProfile))
		if err != nil {
			log.Fatal("Unable to load AWS profile with error:", err)
		}
	}

	return cognitoidentityprovider.NewFromConfig(cfg)
}
