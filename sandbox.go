package main

import (
    "encoding/json"
	"log"
	"os"
)

type Configuration struct {
	AdminUserId string
	DefaultChannelId string
	Token string
	LogFile string
}

var configuration = Configuration{}

func main() {
	ReadConfiguration("config.json")
	log.Println(configuration.LogFile)
	log.Println(configuration.AdminUserId)
	log.Println(configuration.DefaultChannelId)
	log.Println(configuration.Token)
}

// Readies config values
func ReadConfiguration(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening configuration file...?")
	}
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if(err != nil) {
		log.Fatal("Error parsing configuration file")
	}
}