package main

import (
	"fmt"
	"log"

	serviceConfig "github.com/SnackLog/service-config-lib"
	//"github.com/gin-gonic/gin"
)

func main() {
	err := serviceConfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load service configuration: %v", err))
	}

	log.Println(serviceConfig.GetConfig().ServiceName)

	//router := gin.Default()


	
}
