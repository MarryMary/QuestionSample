package Model

import (
	"triela/Structs"
)

func Ex_Push(ExName string, ExYear string, ExSeason string, ExGenre string, ExTags string, ExTimeUp string) bool {
	db := ConnectDB()
	stmt, err := db.Prepare("INSERT INTO Exercise(ExName, ExType, ExYear, ExSeason, ExTags, ExTimeUp) VALUES (?, ?, ?, ?, ?, ?)")
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
		var ExName, ExType, ExYear, ExSeason, ExTags, ExTimeLimit string

		if err := stmt.Scan(&ExName, &ExYear, &ExSeason, &ExType, &ExTags, &ExTimeLimit); err != nil {
			panic(err)
		}

		Exercise := Structs.Exercise{
			ExName:      ExName,
			ExType:      ExType,
			ExYear:      ExYear,
			ExSeason:    ExSeason,
			ExTags:      ExTags,
			ExTimeLimit: ExTimeLimit,
		}

		Exercises = append(Exercises, Exercise)
	}

	if err = stmt.Err(); err != nil {
		panic(err)
	}

	return Exercises
}

func Ex_Find(ExeName string, ExeType string) Structs.Exercise {
	db := ConnectDB()
	stmt := db.QueryRow("SELECT * FROM Exercise WHERE ExName = ? AND ExType = ?", ExeName, ExeType)

	var ExName, ExYear, ExSeason, ExType, ExTags, ExTimeLimit string

	if err := stmt.Scan(&ExName, &ExYear, &ExSeason, &ExType, &ExTags, &ExTimeLimit); err != nil {
		panic(err)
	}
	Exercise := Structs.Exercise{
		ExName:      ExName,
		ExType:      ExType,
		ExYear:      ExYear,
		ExSeason:    ExSeason,
		ExTags:      ExTags,
		ExTimeLimit: ExTimeLimit,
	}

	if err := stmt.Err(); err != nil {
		panic(err)
	}
	return Exercise
}

func Ex_Update(OldExName string, OldExType string, ExName string, ExYear string, ExSeason string, ExType string, ExTags string, ExTimeLimit string) bool {
	db := ConnectDB()
	stmt, err := db.Prepare("UPDATE Exercise SET ExName = ?, ExType = ?, ExYear = ?, ExSeason = ?, ExTags = ?, ExLimit = ? WHERE (ExName = ? AND ExType = ?)")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(ExName, ExType, ExYear, ExSeason, ExTags, ExTimeLimit, OldExName, OldExType)

	if err != nil {
		panic(err)
	}

	return false
}
