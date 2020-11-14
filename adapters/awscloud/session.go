package awscloud

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
)

// Client store aws info
type Client struct {
	Session *session.Session
}

// New init aws client session
func New(profile string) *Client {
	return newWithCredentials(profile, nil)
}

func NewAssumed(profile string, mfaSerial string, mfaCode string) *Client {
	baseClient := New(profile)
	sts := baseClient.StsClient()
	tokenOutput, err := sts.GetSessionToken(mfaSerial, mfaCode)
	if err != nil {
		log.WithFields(log.Fields{
			"tokenCode": mfaCode,
			"error":     err,
		}).Error("Error getting session token")
		os.Exit(1)
	}

	stsCreds := credentials.NewCredentials(&TokenCreds{
		AccessKeyId:     *tokenOutput.Credentials.AccessKeyId,
		SecretAccessKey: *tokenOutput.Credentials.SecretAccessKey,
		SessionToken:    *tokenOutput.Credentials.SessionToken,
	})
	return newWithCredentials(profile, stsCreds)
}

func newWithCredentials(profile string, creds *credentials.Credentials) *Client {
	sessionOptions := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}

	if profile != "" {
		sessionOptions.Profile = profile
	}
	if creds != nil {
		sessionOptions.Config.Credentials = creds
	}

	return &Client{
		Session: session.Must(session.NewSessionWithOptions(sessionOptions)),
	}

}
