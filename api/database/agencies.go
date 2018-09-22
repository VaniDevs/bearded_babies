package database

import (
	"../entity"
	"database/sql"
	"fmt"
)

func Agencies(_range []int, _sort []string) []*entity.Agency {
	db := getDatabase()
	defer db.Close()

	query := "SELECT id, login, password, role, name, phone, email, city, address1, address2, contact FROM agency"
	if len(_sort) >= 2 {
		query = fmt.Sprintf("SELECT id, login, password, role, name, phone, email, city, address1, address2, "+
			"contact FROM agency ORDER BY %s %s",
			_sort[0], _sort[1])
	}
	rows, err := db.Query(query)
	checkErr(err)

	agencies := []*entity.Agency{}
	for rows.Next() {
		agency := readAgency(rows)
		agencies = append(agencies, agency)
	}
	return agencies
}

func AddAgency(agency *entity.Agency) *entity.Agency {
	db := getDatabase()
	defer db.Close()

	var insertId int
	err := db.QueryRow("INSERT INTO agency(login, password, role, name, phone, email, city, address1, address2, contact) VALUES "+
		"($1, md5($2),$3,$4,$5,$6,$7,$8,$9,$10) returning id;", agency.Login, agency.Password, agency.Role, agency.Name,
		agency.Phone, agency.Email, agency.City, agency.Address1, agency.Address2, agency.Contact).Scan(&insertId)
	checkErr(err)
	agency.ID = insertId
	return agency
}

func GetAgency(id int) *entity.Agency {
	db := getDatabase()
	defer db.Close()
	rows, err := db.Query("SELECT id, login, password, role, name, phone, email, city, address1, address2, contact "+
		"FROM agency WHERE id = $1", id)
	checkErr(err)
	var agency *entity.Agency
	if rows.Next() {
		agency = readAgency(rows)
	}
	return agency
}

func UpdateAgency(agency *entity.Agency) {
	db := getDatabase()
	defer db.Close()

	_, err := db.Query("UPDATE agency SET login = $2, role = $3, name = $4, phone = $5, email = $6, city = $7, "+
		"address1 = $8, address2 = $9, contact = $10 WHERE id = $1", agency.ID, agency.Login, agency.Role,
		agency.Name, agency.Phone, agency.Email, agency.City, agency.Address1, agency.Address2, agency.Contact)
	checkErr(err)

	if len(agency.Password) > 0 {
		_, err := db.Query("UPDATE agency SET password = $2 WHERE id = $1", agency.ID, agency.Password)
		checkErr(err)
	}
}

func readAgency(rows *sql.Rows) *entity.Agency {
	var id, role int
	var login, password, name, phone, email, city, address1, address2, contact string

	err := rows.Scan(&id, &login, &password, &role, &name, &phone, &email, &city, &address1, &address2, &contact)
	checkErr(err)
	agency := &entity.Agency{ID: id, Login: login, Password: password, Role: role, Name: name,
		Phone: phone, Email: email, City: city, Address1: address1, Address2: address2, Contact: contact}
	return agency
}
