package database

import (
	"../entity"
	"../notification"
	"database/sql"
	"fmt"
	"time"
)

func Clients(_range []int, _sort []string, _filter string, userId int, role int) []*entity.Client {
	db := getDatabase()
	defer db.Close()

	if role == 1 {
		query := "SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, agency_id, " +
			"unemployed, newcomer, homeless, special_needs FROM client"

		if len(_filter) > 0 {
            query += fmt.Sprintf(" WHERE %s ", _filter)
        }
        if len(_sort) >= 2 {
            query += fmt.Sprintf(" ORDER BY %s %s ", _sort[0], _sort[1])
        }

		rows, err := db.Query(query)
		checkErr(err)

		clients := []*entity.Client{}
		for rows.Next() {
			client := readClient(rows)
			clients = append(clients, client)
		}
		return clients
	} else {
		query := "SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, agency_id, " +
			"unemployed, newcomer, homeless, special_needs FROM client WHERE agency_id = $1"
		if len(_sort) >= 2 {
			query = fmt.Sprintf("SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, "+
				"agency_id, unemployed, newcomer, homeless, special_needs FROM client WHERE agency_id = $1 ORDER BY %s %s",
				_sort[0], _sort[1])
		}

		rows, err := db.Query(query, userId)
		checkErr(err)

		clients := []*entity.Client{}
		for rows.Next() {
			client := readClient(rows)
			clients = append(clients, client)
		}
		return clients
	}
}

func AddClient(client *entity.Client) *entity.Client {
	db := getDatabase()
	defer db.Close()

	var insertId int
	err := db.QueryRow("INSERT INTO client(status, name, dob, childdob, phone, email, city, address1, address2, "+
		"agency_id, unemployed, newcomer, homeless, special_needs) VALUES "+
		"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) returning id;", client.Status, client.Name, client.DOB,
		client.ChildDOB, client.Phone, client.Email, client.City, client.Address1, client.Address2,
		client.AgencyId, client.Unemployed, client.Newcomer, client.Homeless, client.SpecialNeeds).Scan(&insertId)
	checkErr(err)
	client.ID = insertId
	return client
}

func GetClient(id int) *entity.Client {
	db := getDatabase()
	defer db.Close()
	rows, err := db.Query("SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, "+
		"agency_id, unemployed, newcomer, homeless, special_needs FROM client WHERE id = $1", id)
	checkErr(err)
	var client *entity.Client
	if rows.Next() {
		client = readClient(rows)
	}
	return client
}

func UpdateClient(client *entity.Client) {

	savedClient := GetClient(client.ID)

	db := getDatabase()
	defer db.Close()

	_, err := db.Query("UPDATE client SET status = $2, name = $3, dob = $4, childdob = $5, phone = $6, email = $7, "+
		"city = $8, address1 = $9, address2 = $10, agency_id = $11, unemployed = $12, newcomer = $13, "+
		"homeless = $14, special_needs = $15 WHERE id = $1", client.ID, client.Status, client.Name, client.DOB,
		client.ChildDOB, client.Phone, client.Email, client.City, client.Address1, client.Address2,
		client.AgencyId, client.Unemployed, client.Newcomer, client.Homeless, client.SpecialNeeds)
	checkErr(err)

	if savedClient.Status == 0 && client.ID == 1 {
		notification.Send(client.Phone)
	}

}

func readClient(rows *sql.Rows) *entity.Client {
	var id, status, agencyId, unemployed, newcomer, homeless, specialNeeds int
	var DOB, ChildDOB time.Time
	var name, phone, email, city, address1, address2 string

	err := rows.Scan(&id, &status, &name, &DOB, &ChildDOB, &phone, &email, &city, &address1, &address2,
		&agencyId, &unemployed, &newcomer, &homeless, &specialNeeds)
	checkErr(err)
	client := &entity.Client{ID: id, Status: status, Name: name, DOB: DOB, ChildDOB: ChildDOB, Phone: phone, Email: email,
		City: city, Address1: address1, Address2: address2, AgencyId: agencyId,
		Unemployed: unemployed, Newcomer: newcomer, Homeless: homeless, SpecialNeeds: specialNeeds}
	return client
}
