package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InitReferrals(router *gin.Engine) {
	router.GET("/referrals", Referrals)
	router.POST("/referrals", AddReferral)
	router.GET("/referrals/:id", Referral)
	router.PUT("/referrals/:id", PutReferral)
}

func Referrals(c *gin.Context) {
	_, _range, _sort := GetListParams(c)
	referrals := database.Referrals(_range, _sort)
	SetContentRange(c, "referrals", 0, len(referrals), len(referrals))
	c.JSON(http.StatusOK, referrals)
}

func Referral(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, database.GetReferral(id))
}

func AddReferral(c *gin.Context) {
	/*id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")*/
	var referral *entity.Referral
	c.BindJSON(&referral)
	c.JSON(http.StatusOK, database.AddReferral(referral))
}

func PutReferral(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var referral *entity.Referral
	c.BindJSON(&referral)
	database.UpdateReferral(referral)
	c.JSON(http.StatusOK, database.GetReferral(id))
}
