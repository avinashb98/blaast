package application

import (
	"github.com/avinashb98/blaast/config"
	"github.com/avinashb98/blaast/datasources"
	blaastservice "github.com/avinashb98/blaast/domain/blaast"
	"github.com/avinashb98/blaast/http"
	handler "github.com/avinashb98/blaast/http"
	blaastrepo "github.com/avinashb98/blaast/repository/blaast"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.New()

func StartApplication() {
	conf := getConfig()
	initialiseMongo(conf.Mongo)

	router.Use(gin.Recovery())
	router.GET("/status", http.GetStatus)

	blaastRepo := blaastrepo.New(conf.Mongo)
	blaastService := blaastservice.New(blaastRepo)
	blaastHandler := handler.NewBlaast(blaastService)

	blaastV1 := router.Group("/api/v1/blaast")
	{
		blaastV1.POST("/", blaastHandler.CreateBlaast)
		blaastV1.GET("/", blaastHandler.GetByID)
	}

	err := router.Run(":" + conf.Server.Port)
	if err != nil {
		log.Panic(err)
	}
}

func getConfig() *config.Config {
	conf, err := config.GetConfig()
	if err != nil {
		log.Panic(err)
	}
	return conf
}

func initialiseMongo(conf *config.Mongo) {
	err := datasources.InitMongoORM(conf)
	if err != nil {
		log.Panic(err)
	}
}
