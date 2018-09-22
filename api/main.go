package main

import (
	"./service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	service.InitAgencies(router)
	service.InitClients(router)
	service.InitGears(router)
	service.InitReferrals(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"POST", "GET", "OPTION", "PUT"},
		AllowHeaders:    []string{"Content-Type"},
		ExposeHeaders:   []string{"Content-Range"},
		AllowOriginFunc: original,
	}))
	return router
}

func original(origin string) bool {
	//TODO: Only http://localhost:3000
	return true
}
