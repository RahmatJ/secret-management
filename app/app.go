package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"secret-management/internal/di"
)

func Run() {
	envConfig, err := initEnv()
	if err != nil {
		log.Fatal("cannot failed read env")
		return
	}

	//TODO(Rahmat): Add usage envConfig
	fmt.Printf("%+v", envConfig)

	r := gin.New()

	api := r.Group("/api")
	internal := api.Group("/private")

	_, err = di.InitializeDependency(internal)
	if err != nil {
		log.Fatal("cannot init dependency")
		return
	}

	err = r.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
