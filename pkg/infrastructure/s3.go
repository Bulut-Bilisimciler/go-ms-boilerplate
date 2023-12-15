package infrastructure

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/rs/zerolog/log"
)

func NewS3Session(connStr string) *s3.S3 {
	// Parse the URL
	parsed, err := url.Parse(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("app.s3.init: failed to parse s3 connection string")
	}

	// Extract components from the parsed URL
	scheme := parsed.Scheme // e.g. "s3"
	accessKey := parsed.User.Username()
	secretKey, _ := parsed.User.Password()
	host := parsed.Hostname()
	port := parsed.Port()
	region := strings.TrimPrefix(parsed.Path, "/")

	// open s3 session
	creds := credentials.NewStaticCredentials(
		accessKey,
		secretKey,
		"",
	)

	// Disable certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpClient := &http.Client{Transport: tr}

	// create session
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(scheme + "://" + host + ":" + port),
		Region:           aws.String(region),
		Credentials:      creds,
		S3ForcePathStyle: aws.Bool(true),
		HTTPClient:       httpClient,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("app.s3.init: failed to create s3 session")
	}
	// crete session
	s3Client := s3.New(sess)

	return s3Client
}
