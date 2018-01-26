package main

import (
	"uploadS3/server/beanstalk"
	"fmt"
)

func main() {
	fmt.Printf("Starting Upload Files Server\n")
	beanstalk.StartConsumer()
}