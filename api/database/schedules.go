package database

import (
	"../entity"
	"time"
)

func Schedules(date time.Time) []*entity.TimeInterval {
	db := getDatabase()
	defer db.Close()

	bd := bod(date)
	ed := eod(date)

	rows, err := db.Query("SELECT appointment1, appointment2 FROM referral WHERE "+
		"(appointment1 >= $1 AND appointment1 <= $2) OR (appointment2 >= $1 AND appointment2 <= $2)", bd, ed)
	checkErr(err)

	intervals := []*entity.TimeInterval{}
	for rows.Next() {
		var appointment1, appointment2 time.Time
		err := rows.Scan(&appointment1, &appointment2)
		checkErr(err)
		if appointment1.After(bd) && appointment1.Before(ed) {
			interval := &entity.TimeInterval{Time: appointment1, Interval: 30}
			intervals = append(intervals, interval)
		}
		if appointment2.After(bd) && appointment2.Before(ed) {
			interval := &entity.TimeInterval{Time: appointment2, Interval: 15}
			intervals = append(intervals, interval)
		}
	}
	return intervals
}

func bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func eod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}
