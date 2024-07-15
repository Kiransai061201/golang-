package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteBucket(client *s3.S3, bucketName string) error {
	_, err := client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	s3Client := s3.New(sess)

	bucketName := "thisbucketcreatedwithgolang"
	err = DeleteBucket(s3Client, bucketName)
	if err != nil {
		fmt.Printf("Couldn't delete bucket: %v", err)
		return
	}

	fmt.Println("Successfully deleted bucket")
}
