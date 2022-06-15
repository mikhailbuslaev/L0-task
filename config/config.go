package config

import (
	"os"
	"log"
	"gopkg.in/yaml.v3"
)

func Parse(any interface{}, configFile string) {

	buf, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Cannot read config file")
	}

	err = yaml.Unmarshal(buf, any)
	if err != nil {
		log.Println(err)
		log.Fatalf("Cannot parse config")
	}
}

