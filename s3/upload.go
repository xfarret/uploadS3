package s3

import (
	"fmt"
	"uploadS3/models"
	//"time"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"uploadS3/config"
	"os"
)

func UploadFile(id uint64, soft *models.Software) {
	fmt.Printf("Uploading file from queue: %d - %s\n", id, soft)
	//time.Sleep(10*time.Second)
	s3uploadfile()
	fmt.Printf("File from queue %d uploaded\n", id)
}

func s3uploadfile() {
	params := config.GetParameters()

	creds := credentials.Value{
		AccessKeyID: params.AWS.Key,
		SecretAccessKey: params.AWS.Secret,
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(params.AWS.Region),
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	fmt.Printf("uploader: %x", uploader )
	// @todo set dynamic name given in beanstalk server
	filename := "/home/xfarret/Téléchargements/PhpStorm-2017.3.2.tar.gz"
	f, err  := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %q, %v", filename, err)
		return
		//return fmt.Printf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(params.AWS.BucketName),
		Key:    aws.String("phpstorm.tgz"),
		Body:   f,
	})
	if err != nil {
		fmt.Errorf("failed to upload file, %v", err)
		return
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)
}
