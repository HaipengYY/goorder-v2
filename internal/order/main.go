package main

import (
	"github.com/AfRpEng/gorder-v2/common/config"
	"github.com/spf13/viper"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	log.Println("%v", viper.Get("order"))

}
