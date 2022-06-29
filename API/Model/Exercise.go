package Model

func Ex_Push(ExName string, ExDetail string, ExTags string, ExTimeUp string) {
	db := ConnectDB()
	stmt, err := db.Prepare("INSERT INTO Exercise(ExName, ExDetail, ExTags, ExTimeUp) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(ExName, ExDetail, ExTags, ExTimeUp)

	if err != nil {
		panic(err)
	}
}

func Ex_Search() {}
