package database

import (
	"../entity"
	"database/sql"
	"fmt"
)

func Gears(_range []int, _sort []string) []*entity.Gear {
	db := getDatabase()
	defer db.Close()

	query := "SELECT id, name FROM gear"
	if len(_sort) >= 2 {
		query = fmt.Sprintf("SELECT id, name FROM gear ORDER BY %s %s", _sort[0], _sort[1])
	}
	rows, err := db.Query(query)
	checkErr(err)

	gears := []*entity.Gear{}
	for rows.Next() {
		gear := readGear(rows)
		gears = append(gears, gear)
	}
	return gears
}

func AddGear(gear *entity.Gear) *entity.Gear {
	db := getDatabase()
	defer db.Close()

	var insertId int
	err := db.QueryRow("INSERT INTO gear(name) VALUES ($1) returning id;", gear.Name).Scan(&insertId)
	checkErr(err)
	gear.ID = insertId
	return gear
}

func GetGear(id int) *entity.Gear {
	db := getDatabase()
	defer db.Close()
	rows, err := db.Query("SELECT id, name FROM gear WHERE id = $1", id)
	checkErr(err)
	var gear *entity.Gear
	if rows.Next() {
		gear = readGear(rows)
	}
	return gear
}

func UpdateGear(gear *entity.Gear) {
	db := getDatabase()
	defer db.Close()

	_, err := db.Query("UPDATE gear SET name = $2 WHERE id = $1", gear.ID, gear.Name)
	checkErr(err)
}

func readGear(rows *sql.Rows) *entity.Gear {
	var id int
	var name string

	err := rows.Scan(&id, &name)
	checkErr(err)
	gear := &entity.Gear{ID: id, Name: name}
	return gear
}
