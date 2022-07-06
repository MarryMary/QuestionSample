package Model

import (
	"triela/Structs"
)

func Ex_Push(ExName string, ExYear string, ExSeason string, ExGenre string, ExTags string, ExTimeUp string) bool {
	db := ConnectDB()
	stmt, err := db.Prepare("INSERT INTO Exercise(ExName, ExYear, ExSeason, ExType, ExTags, ExTimeUp) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(ExName, ExYear, ExSeason, ExGenre, ExTags, ExTimeUp)

	if err != nil {
		panic(err)
	}

	return false
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
		var ExName, ExYear, ExSeason, ExType, ExTags, ExTimeUp string

		if err := stmt.Scan(&ExId, &ExName, &ExYear, &ExSeason, &ExType, &ExTags, &ExTimeUp); err != nil {
			panic(err)
		}

		Exercise := Structs.Exercise{
			ExId:     ExId,
			ExName:   ExName,
			ExYear:   ExYear,
			ExSeason: ExSeason,
			ExType:   ExType,
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

func Ex_Find(ExeId *string) Structs.Exercise {
	db := ConnectDB()
	stmt := db.QueryRow("SELECT * FROM Exercise WHERE ExId = ?", ExeId)

	var ExId int
	var ExName, ExYear, ExSeason, ExType, ExTags, ExTimeUp string

	if err := stmt.Scan(&ExId, &ExName, &ExYear, &ExSeason, &ExType, &ExTags, &ExTimeUp); err != nil {
		panic(err)
	}
	Exercise := Structs.Exercise{
		ExId:     ExId,
		ExName:   ExName,
		ExYear:   ExYear,
		ExSeason: ExSeason,
		ExType:   ExType,
		ExTags:   ExTags,
		ExTimeUp: ExTimeUp,
	}

	if err := stmt.Err(); err != nil {
		panic(err)
	}
	return Exercise
}

func Ex_Update(ExId string, ExName string, ExYear string, ExSeason string, ExType string, ExTags string, ExTimeUp string) bool {
	db := ConnectDB()
	stmt, err := db.Prepare("UPDATE Exercise SET ExName = ?, ExYear = ?, ExSeason = ?, ExType = ?, ExTags = ?, ExTimeUp = ? WHERE (ExId = ?)")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(ExName, ExYear, ExSeason, ExType, ExTags, ExTimeUp, ExId)

	if err != nil {
		panic(err)
	}

	return false
}
