package awscloud

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type TokenCreds struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
}

func (t *TokenCreds) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     t.AccessKeyId,
		SecretAccessKey: t.SecretAccessKey,
		SessionToken:    t.SessionToken,
	}, nil
}

func (t *TokenCreds) IsExpired() bool {
	return false
}
