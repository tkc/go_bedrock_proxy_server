package server

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"net/http"
	"time"
)

// Sign creates and signs an HTTP request to AWS Bedrock
func sign(payloadBytes []byte, config *Config) (*http.Request, error) {
	// Setting the endpoint and path
	endpoint := fmt.Sprintf("https://bedrock-runtime.%s.amazonaws.com", config.Region)
	path := fmt.Sprintf("/model/%s/invoke", config.ModelID)
	url := endpoint + path
	body := bytes.NewReader(payloadBytes)

	// Creating the HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Setting the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Amz-Target", "BedrockRuntime.InvokeModel")

	// Creating the AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyID,
			config.SecretAccessKey,
			"",
		),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	// Resetting the body reader to reuse for signing
	bodyReader := bytes.NewReader(payloadBytes) // io.ReadSeeker interface

	// Signing the request
	signer := v4.NewSigner(sess.Config.Credentials)
	_, err = signer.Sign(req, bodyReader, "bedrock", config.Region, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	return req, nil
}
