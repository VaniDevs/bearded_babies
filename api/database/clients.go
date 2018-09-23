package database

import (
	"../entity"
	"database/sql"
	"fmt"
	"time"
)

func Clients(_range []int, _sort []string) []*entity.Client {
	db := getDatabase()
	defer db.Close()

	query := "SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, notification, agency_id, " +
		"unemployed, newcomer, homeless, special_needs FROM client"
	if len(_sort) >= 2 {
		query = fmt.Sprintf("SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, "+
			"notification, agency_id, unemployed, newcomer, homeless, special_needs FROM client ORDER BY %s %s",
			_sort[0], _sort[1])
	}
	rows, err := db.Query(query)
	checkErr(err)

	clients := []*entity.Client{}
	for rows.Next() {
		client := readClient(rows)
		clients = append(clients, client)
	}
	return clients
}

func AddClient(client *entity.Client) *entity.Client {
	db := getDatabase()
	defer db.Close()

	var insertId int
	err := db.QueryRow("INSERT INTO client(status, name, dob, childdob, phone, email, city, address1, address2, "+
		"notification, agency_id, unemployed, newcomer, homeless, special_needs) VALUES "+
		"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) returning id;", client.Status, client.Name, client.DOB,
		client.ChildDOB, client.Phone, client.Email, client.City, client.Address1, client.Address2, client.Notification,
		client.AgencyId, client.Unemployed, client.Newcomer, client.Homeless, client.SpecialNeeds).Scan(&insertId)
	checkErr(err)
	client.ID = insertId
	return client
}

func GetClient(id int) *entity.Client {
	db := getDatabase()
	defer db.Close()
	rows, err := db.Query("SELECT id, status, name, dob, childdob, phone, email, city, address1, address2, notification, "+
		"agency_id, unemployed, newcomer, homeless, special_needs FROM client WHERE id = $1", id)
	checkErr(err)
	var client *entity.Client
	if rows.Next() {
		client = readClient(rows)
	}
	return client
}

func UpdateClient(client *entity.Client) {
	db := getDatabase()
	defer db.Close()

	_, err := db.Query("UPDATE client SET status = $2, name = $3, dob = $4, childdob = $5, phone = $6, email = $7, "+
		"city = $8, address1 = $9, address2 = $10, notification = $11, agency_id = $12, unemployed = $13, newcomer = $14, "+
		"homeless = $15, special_needs = $16 WHERE id = $1", client.ID, client.Status, client.Name, client.DOB,
		client.ChildDOB, client.Phone, client.Email, client.City, client.Address1, client.Address2, client.Notification,
		client.AgencyId, client.Unemployed, client.Newcomer, client.Homeless, client.SpecialNeeds)
	checkErr(err)

}

func readClient(rows *sql.Rows) *entity.Client {
	var id, status, notification, agencyId, unemployed, newcomer, homeless, specialNeeds int
	var DOB, ChildDOB time.Time
	var name, phone, email, city, address1, address2 string

	err := rows.Scan(&id, &status, &name, &DOB, &ChildDOB, &phone, &email, &city, &address1, &address2, &notification,
		&agencyId, &unemployed, &newcomer, &homeless, &specialNeeds)
	checkErr(err)
	client := &entity.Client{ID: id, Status: status, Name: name, DOB: DOB, ChildDOB: ChildDOB, Phone: phone, Email: email,
		City: city, Address1: address1, Address2: address2, Notification: notification, AgencyId: agencyId,
		Unemployed: unemployed, Newcomer: newcomer, Homeless: homeless, SpecialNeeds: specialNeeds}
	return client
}
