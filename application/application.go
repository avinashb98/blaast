package application

import (
	"github.com/avinashb98/blaast/config"
	"github.com/avinashb98/blaast/datasources"
	"github.com/avinashb98/blaast/http"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.New()

func StartApplication() {
	router.Use(gin.Recovery())
	conf, err := config.GetConfig()
	if err != nil {
		log.Panic(err)
	}
	err = datasources.InitMongoORM(conf.Mongo)
	if err != nil {
		log.Panic(err)
	}

	router.GET("/status", http.GetStatus)
	err = router.Run(":" + conf.Server.Port)
	if err != nil {
		log.Panic(err)
	}
}
