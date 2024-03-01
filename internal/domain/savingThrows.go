package domain

type SavingThrow struct {
	SavingThrowId int `json:"saving_thow_id"`
	ClassId       int `json:"class_id"`
	Str           int `json:"str"`
	Dex           int `json:"dex"`
	Int           int `json:"int"`
	Con           int `json:"con"`
	Wiz           int `json:"wiz"`
	Cha           int `json:"cha"`
}
