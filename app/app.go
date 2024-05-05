package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"secret-management/config/database"
	"secret-management/internal/di"
)

func Run() {
	envConfig, err := GetEnv()
	if err != nil {
		log.Fatal(fmt.Sprintf("cannot failed read env. error: %+v", err))
		return
	}

	conn, err := database.NewPostgresqlDatabase(&envConfig)
	if err != nil {
		log.Fatal(fmt.Sprintf("error when connection db: %+v", err))
		return
	}

	sql, err := conn.DB()
	if err != nil {
		log.Fatal(fmt.Sprintf(err.Error()))
		return
	}
	defer func() {
		errClose := sql.Close()
		if errClose != nil {
			log.Fatal(errClose.Error())
		}
	}()

	r := gin.New()

	api := r.Group("/api")
	internal := api.Group("/private")

	handler, err := di.InitializeDependency(internal, conn, envConfig)
	if err != nil {
		log.Fatal("cannot init dependency")
		return
	}

	//Start Cron
	handler.CronSetup.Initiate()

	err = r.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
