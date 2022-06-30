package Model

import (
	"triela/Structs"
)

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

func Ex_GetAll() []Structs.Exercise {
	db := ConnectDB()
	stmt, err := db.Query("SELECT * FROM Exercise")
	if err != nil {
		panic(err)
	}

	Exercises := []Structs.Exercise{}

	for stmt.Next() {
		var ExId int
		var ExName, ExDetail, ExTags, ExTimeUp string

		if err := stmt.Scan(&ExId, &ExName, &ExDetail, &ExTags, &ExTimeUp); err != nil {
			panic(err)
		}

		Exercise := Structs.Exercise{
			ExId:     ExId,
			ExName:   ExName,
			ExDetail: ExDetail,
			ExTags:   ExTags,
			ExTimeUp: ExTimeUp,
		}

		Exercises = append(Exercises, Exercise)
	}

	if err = stmt.Err(); err != nil {
		panic(err)
	}

	return Exercises
}
