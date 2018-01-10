package service

import (
	"flag"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)


func SetupTestConfig() {

	var (
		conf = flag.String("conf", "./../../config/config", "config file path")
	)

	if *conf == "" {
		*conf = "./../config/prod"
	}

	flag.Parse()

	fmt.Println(*conf)
	viper.SetConfigName(*conf)
	viper.AddConfigPath("./")
	viper.SetConfigType("json")
	log.SetLevel(log.InfoLevel)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:", err.Error())
		fmt.Println("Read config fail:", err.Error())
	}

}


