package beanstalk

import (
	"github.com/kr/beanstalk"
	"uploadS3/config"
	"fmt"
	"log"
	"time"
	"encoding/json"
	"uploadS3/s3"
	"uploadS3/models"
)

func dial(verbose bool) *beanstalk.Conn {
	params := config.GetParameters()
	if verbose == true {
		fmt.Printf("Connection to Beanstalk\n")
	}
	c, err := beanstalk.Dial("tcp", fmt.Sprintf("%s:%s", params.Beanstalk.Address, params.Beanstalk.Port))
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func StartConsumer() {
	c := dial(true)
	fmt.Printf("Waiting for jobs\n")
	for {
		id, body, err := c.Reserve(120*time.Second)

		if err == nil {
			var soft *models.Software
			json.Unmarshal(body, &soft)

			go s3.UploadFile(id, soft)

			c.Delete(id)
		}
	}
}

func Put(b []byte) (uint64, error) {
	c := dial(false)
	id, err := c.Put(b, 1, 0, 120*time.Second)
	c.Close()
	return id, err
}