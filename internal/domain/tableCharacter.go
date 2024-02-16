package domain

type TableCharacter struct{
	Idcharacter int64 `json:"idcharacter"`
	UserId string `json:"userid"`
	Name string `json:"name"`
	ClassId int `json:"classid"`
	RaceId int `json:"raceid"`
	Background string `json:"background"`
	Hitpoints int `json:"hitpoints"`
	Speed int `json:"speed"`
	ArmorClass int `json:"armorclass"`
	Level int `json:"level"`
	Exp int `json:"exp"`
	CampaignId int `json:"campaignid"`
	Str int `json:"str"`
	Dex int `json:"dex"`
	Int int `json:"int"`
	Wiz int `json:"wiz"`
	Con int `json:"con"`
	Cha int `json:"cha"`
}