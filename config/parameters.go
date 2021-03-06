package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

//type param struct {
//	Parameters map[string][]string
//}


type Parameters struct {
	Beanstalk   struct {
		Address     string `yaml:"address"`
		Port 		string `yaml:"port"`

	} `yaml:"beanstalk"`
	RestServer  struct {
		Address     string `yaml:"address"`
		Port 		string `yaml:"port"`

	} `yaml:"restserver"`
	AWS 		struct {
		BucketName string `yaml:"bucket"`
		Key        string `yaml:"key"`
		Secret     string `yaml:"secret"`
		Region     string `yaml:"region"`
	} `yaml:"aws"`
}


var c *Parameters

//func init() {
//	yamlFile, err := ioutil.ReadFile("parameters.yml")
//	if err != nil {
//		log.Fatalf("yamlFile.Get err   #%v ", err)
//	}
//	err = yaml.Unmarshal(yamlFile, config)
//	if err != nil {
//		log.Fatalf("Unmarshal: %v", err)
//	}
//
//	//database = make(map[string]interface{})
//	//
//	//beanstalk := make([]string, 2)
//	//
//	//print(beanstalk)
//	//
//	//restserver := make([]string, 2)
//	//print(restserver)
//
//	//brands[0] = Brand{Name: "Gucci"}
//	//brands[1] = Brand{Name: "LV"}
//}


func init() {
	yamlFile, err := ioutil.ReadFile("parameters.yml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func GetParameters() *Parameters {
	return c
}
