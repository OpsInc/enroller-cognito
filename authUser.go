package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (conf *Config) Signin() string {
	InitiateAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(conf.AppClientID),
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,

		AuthParameters: map[string]string{
			"USERNAME": conf.Username,
			"PASSWORD": conf.Password,
		},
	}

	out, err := conf.Cognito.InitiateAuth(context.TODO(), InitiateAuthInput)
	if err != nil {
		log.Fatal("Cognito signin failed because of error: ", err)
	}

	return *out.AuthenticationResult.IdToken
}
