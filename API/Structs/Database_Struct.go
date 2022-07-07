package Structs

import "time"

type Exercise struct {
	ExName      string
	ExType      string
	ExYear      string
	ExSeason    string
	ExTags      string
	ExTimeLimit string
}

type Question struct {
	ExName     string
	ExType     string
	QName      string
	QTags      string
	QMain      string
	AnswerArea string
	QAnswer    string
	QChoice    string
	QScore     int
}

type Respondent struct {
	ManageId   int
	UserId     int
	ExName     string
	ExType     string
	QName      string
	Score      int
	CrWr       int
	RecordedAt time.Time
}

type Fixer struct {
	Sys_Manage_Id int
	Admin_Id      int
	ExName        string
	ExType        string
	QName         string
	Edit_Limit    time.Time
}

type Fix struct {
	Id        int
	ExName    string
	ExType    string
	QName     string
	main_text string
}

type Creator struct {
	ExName  string
	ExType  string
	QName   string
	UserId  int
	Gr_OR_U int
	Q_OR_Ex int
}

type PreUser struct {
	id            int
	user_token    string
	email         string
	register_at   time.Time
	register_type int
	affect_id     int
}

type User struct {
	id              int
	email           string
	user_name       string
	pass            string
	user_pict       string
	GAuthID         string
	delete_at       time.Time
	IsTwoFactor     int
	TwoFactorSecret string
	delete_flag     int
}

type UserRoll struct {
	Management_SysId int
	UserId           int
	Roll             int
}

type RollType struct {
	System_Manage_Id int
	RollType         int
	RollName         string
	RollActivity     string
}
