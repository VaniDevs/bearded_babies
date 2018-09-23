package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Appointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	referral := database.GetReferral(id)
	client := database.GetClient(referral.ClientID)
	c.JSON(http.StatusOK, map[string]interface{}{
	    "id": referral.ID,
        "name": client.Name,
    })
}

func PutAppointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var referral *entity.Referral
	c.BindJSON(&referral)
	database.UpdateReferral(referral)
	c.JSON(http.StatusOK, database.GetReferral(id))
}
