package main

import (
	"github.com/spf13/viper"
	"log"
)

func main(){
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
	}
	log.Printf("uri is %s", config)
}
