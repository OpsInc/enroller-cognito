package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/google/go-github/v57/github"
	"golang.org/x/crypto/nacl/box"
)

const (
	repoName string = "enroller-cognito"
)

// Code copied from:
// https://zostay.com/posts/2022/05/04/do-not-use-libsodium-with-go/
// This allows us to use a public to encrypt a scecret
func encryptSecret(pk, secret string) (string, error) {
	var pkBytes [32]byte
	copy(pkBytes[:], pk)
	secretBytes := []byte(secret)

	out := make([]byte, 0,
		len(secretBytes)+
			box.Overhead+
			len(pkBytes))

	enc, err := box.SealAnonymous(
		out, secretBytes, &pkBytes, rand.Reader,
	)
	if err != nil {
		return "", err
	}

	encEnc := base64.StdEncoding.EncodeToString(enc)

	return encEnc, nil
}

func (conf *GithubConf) UpdateSecrets() int {
	// Connection to Github using the GITHUB_TOKEN env variable
	client := github.NewClientWithEnvProxy().WithAuthToken(conf.GithubToken)

	// Fetch the Github public key from our Repo
	publicKey, _, err := client.Actions.GetRepoPublicKey(context.TODO(), conf.Organization, conf.Repo)
	if err != nil {
		log.Fatalf("Fetchin the github repo: %s PublicKey has failed with error: %v", conf.Repo, err)
	}

	// Encrypt the CognitoToken using the github public key
	encryptedSecret, err := encryptSecret(*publicKey.KeyID, conf.CognitoToken)
	if err != nil {
		log.Fatal("Secret encryption for github as failed with error: ", err)
	}

	// Construct the values required for CreateOrUpdateRepoSecret
	secretConf := &github.EncryptedSecret{
		Name:           "COGNITO_TOKEN",
		EncryptedValue: encryptedSecret,
		KeyID:          *publicKey.KeyID,
	}

	resp, err := client.Actions.CreateOrUpdateRepoSecret(context.TODO(), "Opsinc", repoName, secretConf)
	if err != nil {
		log.Fatal("Writing secret to github as failed with error: ", err)
	}
	log.Printf("Secret %s has been successfully writen to the repo %s", secretConf.Name, conf.Repo)

	return resp.StatusCode
}
