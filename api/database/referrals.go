package database

import (
	"../entity"
	"database/sql"
	"fmt"
	"time"
)

func Referrals(_range []int, _sort []string, userId int, role int) []*entity.Referral {
	db := getDatabase()
	defer db.Close()

	if role == 1 {
		query := "SELECT id, client_id, appointment1, appointment2 FROM referral"
		if len(_sort) >= 2 {
			query = fmt.Sprintf("SELECT id, client_id, appointment1, appointment2 FROM referral ORDER BY %s %s",
				_sort[0], _sort[1])
		}
		rows, err := db.Query(query)
		checkErr(err)

		referrals := []*entity.Referral{}
		for rows.Next() {
			referral := readReferral(rows)
			referrals = append(referrals, referral)
		}
		return referrals
	} else {
		query := "SELECT r.id, client_id, appointment1, appointment2 FROM referral as r INNER JOIN client as c ON c.id = r.client_id WHERE c.agency_id = $1"
		if len(_sort) >= 2 {
			query = fmt.Sprintf("SELECT r.id, client_id, appointment1, appointment2 FROM referral as r INNER JOIN client as c ON c.id = r.client_id WHERE c.agency_id = $1 ORDER BY %s %s",
				_sort[0], _sort[1])
		}
		rows, err := db.Query(query, userId)
		checkErr(err)

		referrals := []*entity.Referral{}
		for rows.Next() {
			referral := readReferral(rows)
			referrals = append(referrals, referral)
		}
		return referrals
	}
}

func AddReferral(referral *entity.Referral) *entity.Referral {
	db := getDatabase()
	defer db.Close()
    t := time.Date(1, 1, 0, 0, 0, 0, 0, time.UTC)
	var insertId int
	err := db.QueryRow("INSERT INTO referral (client_id, appointment1, appointment2) VALUES "+
		"($1,$2,$3) returning id;", referral.ClientID, t, t).Scan(&insertId)
	checkErr(err)
	referral.ID = insertId

	addReferralGear(db, referral)

	return referral
}

func GetReferral(id int) *entity.Referral {
	db := getDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT id, client_id, appointment1, appointment2 FROM referral WHERE id = $1", id)
	checkErr(err)

	var referral *entity.Referral
	if rows.Next() {
		referral = readReferral(rows)
	}
	rows, err = db.Query("SELECT gear_id, status FROM referral_gear WHERE referral_id = $1", id)
	checkErr(err)

	referral.Requested = []int{}
	referral.Unavailable = []int{}

	for rows.Next() {
		var gearId, status int
		err := rows.Scan(&gearId, &status)
		checkErr(err)

		if status == 1 {
			referral.Requested = append(referral.Requested, gearId)
		} else if status == 2 {
			referral.Unavailable = append(referral.Unavailable, gearId)
		}
	}

	return referral
}

func UpdateReferral(referral *entity.Referral) {
	db := getDatabase()
	defer db.Close()

	_, err := db.Query("UPDATE referral SET client_id = $2, appointment1 = $3, appointment2 = $4 WHERE id = $1",
		referral.ID, referral.ClientID, referral.Appointment1, referral.Appointment2)
	checkErr(err)

	_, err = db.Query("DELETE FROM referral_gear WHERE referral_id = $1", referral.ID)
	checkErr(err)

	addReferralGear(db, referral)
}

func readReferral(rows *sql.Rows) *entity.Referral {
	var id, client_id int
	var appointment1, appointment2 time.Time

	err := rows.Scan(&id, &client_id, &appointment1, &appointment2)
	checkErr(err)
	referral := &entity.Referral{ID: id, ClientID: client_id, Appointment1: appointment1, Appointment2: appointment2}
	return referral
}

func addReferralGear(db *sql.DB, referral *entity.Referral) {
	for _, gear := range referral.Requested {
		_, err := db.Exec("INSERT INTO referral_gear (referral_id, gear_id, status) VALUES ($1, $2, $3)",
			referral.ID, gear, 1)
		checkErr(err)
	}
	for _, gear := range referral.Unavailable {
		_, err := db.Exec("INSERT INTO referral_gear (referral_id, gear_id, status) VALUES ($1, $2, $3)",
			referral.ID, gear, 2)
		checkErr(err)
	}
}
