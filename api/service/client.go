package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InitClients(router *gin.Engine) {
	router.GET("/clients", Clients)
	router.POST("/clients", AddClient)
	router.GET("/clients/:id", Clients)
	router.PUT("/clients/:id", PutClient)
}

func Clients(c *gin.Context) {
	_, _range, _sort := GetListParams(c)
	clients := database.Clients(_range, _sort)
	SetContentRange(c, "clients", 0, len(clients), len(clients))
	c.JSON(http.StatusOK, clients)
}

func Client(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, database.GetClient(id))
}

func AddClient(c *gin.Context) {
	/*id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")*/
	var client *entity.Client
	c.BindJSON(&client)
	c.JSON(http.StatusOK, database.AddClient(client))
}

func PutClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var client *entity.Client
	c.BindJSON(&client)
	database.UpdateClient(client)
	c.JSON(http.StatusOK, database.GetClient(id))
}
