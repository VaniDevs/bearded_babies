package service

import (
	"../database"
	"../entity"
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Clients(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	role := int(claims["role"].(float64))
	_filter, _range, _sort := GetListParams(c)
	clients := database.Clients(_range, _sort, _filter, id, role)
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
	fmt.Println(client)
	c.JSON(http.StatusOK, database.AddClient(client))
}

func PutClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var client *entity.Client
	c.BindJSON(&client)
	database.UpdateClient(client)
	c.JSON(http.StatusOK, database.GetClient(id))
}
