package main


func main() {
	go launchBeanstalk()
	go launchRestServer()
}
