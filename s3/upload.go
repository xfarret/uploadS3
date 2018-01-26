package s3

import (
	"fmt"
	"uploadS3/models"
	"time"
)

func UploadFile(id uint64, soft *models.Software) {
	fmt.Printf("Uploading file from queue: %d - %s\n", id, soft)
	time.Sleep(10*time.Second)
	fmt.Printf("File from queue %d uploaded\n", id)
}