package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Agencies(c *gin.Context) {
	_filter, _range, _sort := GetListParams(c)
	agencies := database.Agencies(_range, _sort, _filter)
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
