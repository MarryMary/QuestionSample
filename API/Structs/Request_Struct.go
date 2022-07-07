package Structs

type ExercisePushRequest struct {
	MainTitle string `json:"maintitle"`
	Year      string `json:"year"`
	Season    string `json:"season"`
	Genre     string `json:"genre"`
	Tag       string `json:"tag"`
	Limit     string `json:"limit"`
}

type ExerciseUpdateRequest struct {
	OldTitle  string `json:"oldtitle"`
	OldType   string `json:"oldtype"`
	MainTitle string `json:"maintitle"`
	Year      string `json:"year"`
	Season    string `json:"season"`
	Type      string `json:"type"`
	Tag       string `json:"tag"`
	Limit     string `json:"limit"`
}

type QuestionPushRequest struct {
	ExName    string `json:"exname"`
	ExType    string `json:"extype"`
	MainTitle string `json:"maintitle"`
	Tag       string `json:"tag"`
	MainText  string `json:"maintext"`
	QType     string `json:"qtype"`
	Answer    string `json:"answer"`
	Choice    string `json:"choice"`
	Score     string `json:"score"`
}

type QuestionUpdateRequest struct {
	Id        string `json:"id"`
	MainTitle string `json:"maintitle"`
	Tag       string `json:"tag"`
	MainText  string `json:"maintext"`
	QType     string `json:"qtype"`
	Answer    string `json:"answer"`
	Choice    string `json:"choice"`
	Score     string `json:"score"`
}

type ExerciseFindRequest struct {
	ExName string `json:"exname"`
	ExType string `json:"extype"`
}

type QuestionFindRequest struct {
	ExName string `json:"exname"`
	ExType string `json:"extype"`
	QName  string `json:"qname"`
}
