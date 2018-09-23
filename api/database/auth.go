package database

func User(login string, password string) (int, int) {
	db := getDatabase()
	defer db.Close()

	query := "SELECT id, role FROM agency WHERE login = $1 and password = md5($2)"
	rows, err := db.Query(query, login, password)
	checkErr(err)

	if rows.Next() {
		var id, role int

		err := rows.Scan(&id, &role)
		checkErr(err)
		return id, role
	}
	return 0, 0
}
