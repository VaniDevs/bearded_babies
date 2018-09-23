package service

import (
	"../database"
	"../entity"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Referrals(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	role := int(claims["role"].(float64))
	_, _range, _sort := GetListParams(c)
	referrals := database.Referrals(_range, _sort, id, role)
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
