package service

import (
	"../database"
	"../entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitSchedule(router *gin.Engine) {
	router.GET("/schedules", Schedules)
}

func Schedules(c *gin.Context) {
	t := time.Date(2018, 9, 1, 0, 0, 0, 0, time.UTC)
	intervals := allIntervals(t, 30)
	appointments := database.Schedules(t)
	intervals = exclude(intervals, appointments)
	c.JSON(http.StatusOK, intervals)
}

func startOperations(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 9, 0, 0, 0, date.Location())
}

func endOperations(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 17, 0, 0, 0, date.Location())
}

func allIntervals(date time.Time, interval int) []*entity.TimeInterval {
	intervals := []*entity.TimeInterval{}
	start := startOperations(date)
	end := endOperations(date).Add(-time.Minute * time.Duration(interval)).Add(1)
	dt := start
	for dt.Before(end) {
		interval := &entity.TimeInterval{Time: dt, Interval: interval}
		intervals = append(intervals, interval)
		dt = dt.Add(time.Minute * time.Duration(15))
	}
	return intervals
}

func exclude(intervals []*entity.TimeInterval, appointments []*entity.TimeInterval) []*entity.TimeInterval {
	result := []*entity.TimeInterval{}
	for _, interval := range intervals {
		ist := interval.Time
		iet := interval.Time.Add(time.Minute * time.Duration(interval.Interval))
		busy := false
		for _, appointment := range appointments {
			ast := appointment.Time
			aet := appointment.Time.Add(time.Minute * time.Duration(interval.Interval))
			if ist.Add(-1).Before(ast) && iet.Add(1).After(ast) {
				busy = true
			}
			if ist.Add(-1).After(aet) && iet.Add(1).Before(aet) {
				busy = true
			}
			if ast.Add(1).After(ist) && aet.Add(-1).Before(iet) {
				busy = true
			}
		}
		if !busy {
			result = append(result, interval)
		}
	}
	return result
}
