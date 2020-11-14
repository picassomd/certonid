package awscloud

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
)

type STSClient struct {
	Client *sts.STS
}

func (cl *STSClient) GetSessionToken(mfaSerial string, tokenCode string) (*sts.GetSessionTokenOutput, error) {
	result, err := cl.Client.GetSessionToken(&sts.GetSessionTokenInput{
		SerialNumber: aws.String(mfaSerial),
		TokenCode:    aws.String(tokenCode),
	})
	if err != nil {
		return nil, fmt.Errorf("Error getting session token: %w", err)
	}
	return result, nil
}

func (client *Client) StsClient() *STSClient {
	awsConfig := aws.Config{}

	return &STSClient{
		Client: sts.New(client.Session, &awsConfig),
	}
}
