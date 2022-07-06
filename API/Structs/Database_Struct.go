package Structs

import "time"

/*
エクササイズテーブル（試験グループ用テーブル）
エクササイズ <- この情報を格納
 |
 |--問1
 |   |--小問1
 |
 |--問2

エクササイズグループに問題が属す

ExId -- エクササイズ管理用番号（AutoIncrement）
ExName -- エクササイズ名
ExDetail -- エクササイズ説明文
ExTags -- エクササイズのタグ(カンマ区切り、サーバーで必要に応じて分割)
ExTimeUp -- エクササイズの制限時間
*/
type Exercise struct {
	ExId     int
	ExName   string
	ExYear   string
	ExSeason string
	ExType   string
	ExTags   string
	ExTimeUp string
}

/*
問題テーブル
QId -- 問題管理用番号（AutoIncrement）
QName -- 問題名（問1とか大問1とか）
QMain -- 問題文
QType -- 記述問題か選択問題か(0 or 1)
AnswerArea -- 選択問題なら選択肢を、記述問題ならテキストエリアであるとわかる記号を([SELECT>ア:Example,イ:Sample,ウ:Test,エ:Job]とか[WRITING>TEXTAREA]とか)
QAnswer -- 答えまたは記述の場合は採点ポイントをタグの形式（カンマ区切り）で保存
IncExId -- この問題が属するエクササイズの管理番号
*/

type Question struct {
	QId        int
	QName      string
	QTags      string
	QMain      string
	AnswerArea string
	QAnswer    string
	QChoice    string
	QScore     int
	IncExId    int
}

/*
回答者テーブル
ManageId -- 回答レコードの管理用番号（AutoIncrement）
UId --ユーザー管理番号（ユーザーテーブルについてはログイン画面のGitHub(https://github.com/RewKotoha/AuthSample)を参照）
QuestionId -- 問題テーブル内レコードの管理番号
Score -- 回答のスコア
CrWr -- 正誤判定（正を0、誤を1として記録）
RecordedAt -- 回答日時
*/

type Respondent struct {
	ManageId   int
	UId        int
	QuestionId int
	Score      int
	CrWr       int
	RecordedAt time.Time
}
