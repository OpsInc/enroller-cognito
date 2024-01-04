package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (cognitoConfig *Config) UserSignin(cognitoProvider *cognitoidentityprovider.Client) string {
	//nolint:exhaustruct
	// Ignoring unsused fields: AnalyticsMetadata, ClientMetadata, UserContextData
	InitiateAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(cognitoConfig.AppClientID),
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,

		AuthParameters: map[string]string{
			"USERNAME": cognitoConfig.Username,
			"PASSWORD": cognitoConfig.Password,
		},
	}

	out, err := cognitoProvider.InitiateAuth(context.TODO(), InitiateAuthInput)
	if err != nil {
		log.Fatal("Cognito signin failed because of error: ", err)
	}

	return *out.AuthenticationResult.IdToken
}
