package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		ec2metadataSession := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		}))
		// Use the chain to define multiple sources for credentials
		creds := credentials.NewChainCredentials(
			[]credentials.Provider{
				&credentials.EnvProvider{},
				&ec2rolecreds.EC2RoleProvider{
					Client: ec2metadata.New(ec2metadataSession),
				},
			})

		s3Session := session.Must(session.NewSession(&aws.Config{
			Credentials: creds,
			Region:      aws.String("us-east-1"),
		}))
		svc := s3.New(s3Session)
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("com.sh3r4rd.uploads"),
			Key:    aws.String("BootStrap-script.txt"),
		})

		url, err := req.Presign(15 * time.Minute)
		if err != nil {
			log.Println("Failed to sign request", err)
		}

		log.Println("The URL is", url)
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
