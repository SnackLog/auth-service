package main

import (
	"log"

	serviceConfig "github.com/SnackLog/service-config-lib"
)

func main() {
	serviceConfig.LoadConfig()
	log.Println(serviceConfig.GetConfig().ServiceName)
}
