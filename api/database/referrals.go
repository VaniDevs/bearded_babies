package database

import (
	"../entity"
	"database/sql"
	"fmt"
)

func Referrals(_range []int, _sort []string) []*entity.Referral {
	db := getDatabase()
	defer db.Close()

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
}

func AddReferral(referral *entity.Referral) *entity.Referral {
	db := getDatabase()
	defer db.Close()

	var insertId int
	err := db.QueryRow("INSERT INTO agency(client_id, appointment1, appointment2) VALUES "+
		"($1,$2,$3) returning id;", referral.ClientID, referral.Appointment1, referral.Appointment2).Scan(&insertId)
	checkErr(err)
	referral.ID = insertId
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
	return referral
}

func UpdateReferral(referral *entity.Referral) {
	db := getDatabase()
	defer db.Close()

	_, err := db.Query("UPDATE referral SET client_id = $2, appointment1 = $3, appointment2 = $4 WHERE id = $1",
		referral.ID, referral.ClientID, referral.Appointment1, referral.Appointment2)
	checkErr(err)
}

func readReferral(rows *sql.Rows) *entity.Referral {
	var id, client_id int
	var appointment1, appointment2 entity.NullTime

	err := rows.Scan(&id, &client_id, &appointment1, &appointment2)
	checkErr(err)
	referral := &entity.Referral{ID: id, ClientID: client_id, Appointment1: appointment1, Appointment2: appointment2}
	return referral
}
