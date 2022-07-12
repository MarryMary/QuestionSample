package Model

import (
	"triela/Structs"
)

func Ex_Push(ExName string, ExYear string, ExSeason string, ExGenre string, ExTags string, ExTimeUp string, Gr_OR_U string, UserId any) bool {
	db := ConnectDB()
	stmt, err := db.Prepare("INSERT INTO Exercise(ExName, ExType, ExYear, ExSeason, ExTags, ExTimeUp) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(ExName, ExGenre, ExYear, ExSeason, ExTags, ExTimeUp)

	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("INSERT INTO ExCreator(ExName, ExType, UserId, Gr_OR_U) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(ExName, ExGenre, UserId, Gr_OR_U)

	if err != nil {
		panic(err)
	}

	return false
}

func Ex_GetAll(UserId any, Gr_OR_U string) []Structs.Exercise {
	db := ConnectDB()
	stmt, err := db.Query("SELECT * FROM Exercise LEFT OUTER JOIN ExCreator ON Exercise.ExName=ExCreator.ExName AND Exercise.ExType=ExCreator.ExType WHERE UserId = ? AND Gr_OR_U = ?", UserId, Gr_OR_U)

	if err != nil {
		panic(err)
	}

	Exercises := []Structs.Exercise{}

	for stmt.Next() {
		var ExName, ExType, ExYear, ExSeason, ExTags, ExTimeLimit, Trash, Trasher string
		var UserId, Gr_OR_U int

		if err := stmt.Scan(&ExName, &ExYear, &ExSeason, &ExType, &ExTags, &ExTimeLimit, &Trash, &Trasher, &UserId, &Gr_OR_U); err != nil {
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
