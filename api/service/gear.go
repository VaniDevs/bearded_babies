package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Gears(c *gin.Context) {
	_, _range, _sort := GetListParams(c)
	gears := database.Gears(_range, _sort)
	SetContentRange(c, "gears", 0, len(gears), len(gears))
	c.JSON(http.StatusOK, gears)
}

func Gear(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, database.GetGear(id))
}

func AddGear(c *gin.Context) {
	/*id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")*/
	var gear *entity.Gear
	c.BindJSON(&gear)
	c.JSON(http.StatusOK, database.AddGear(gear))
}

func PutGear(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var gear *entity.Gear
	c.BindJSON(&gear)
	database.UpdateGear(gear)
	c.JSON(http.StatusOK, database.GetGear(id))
}
