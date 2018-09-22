package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InitAgencies(router *gin.Engine) {
	router.GET("/agencies", Agencies)
	router.POST("/agencies", AddAgency)
	router.GET("/agencies/:id", Agency)
	router.PUT("/agencies/:id", PutAgency)
}

func Agencies(c *gin.Context) {
	_, _range, _sort := GetListParams(c)
	agencies := database.Agencies(_range, _sort)
	SetContentRange(c, "agencies", 0, len(agencies), len(agencies))
	c.JSON(http.StatusOK, agencies)
}

func Agency(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, database.GetAgency(id))
}

func AddAgency(c *gin.Context) {
	/*id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")*/
	var agency *entity.Agency
	c.BindJSON(&agency)
	c.JSON(http.StatusOK, database.AddAgency(agency))
}

func PutAgency(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var agency *entity.Agency
	c.BindJSON(&agency)
	database.UpdateAgency(agency)
	c.JSON(http.StatusOK, database.GetAgency(id))
}
