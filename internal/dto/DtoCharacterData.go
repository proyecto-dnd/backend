package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type FullCharacterData struct {
	Character_Id int64  `json:"characterid"`
	User_Id      string `json:"userid"`
	Campaign_Id  int    `json:"campaignid"`
	Race        domain.Race    `json:"race"`
	Class       domain.Class    `json:"class"`
	Background  domain.Background    `json:"background"`
	Name        string `json:"name"`
	Story       string `json:"story"`
	Alignment   string `json:"alignment"`
	Age         int    `json:"age"`
	Hair        string `json:"hair"`
	Eyes        string `json:"eyes"`
	Skin        string `json:"skin"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	ImgUrl      string `json:"img"`
	Str         int    `json:"str"`
	Dex         int    `json:"dex"`
	Int         int    `json:"int"`
	Con         int    `json:"con"`
	Wiz         int    `json:"wiz"`
	Cha         int    `json:"cha"`
	Hitpoints   int    `json:"hitpoints"`
	HitDice     string `json:"hitDice"`
	Speed       int    `json:"speed"`
	Armor_Class  int    `json:"armorclass"`
	Level       int    `json:"level"`
	Exp         int    `json:"exp"`
	Items        []domain.ItemXCharacterData `json:"items"`
	Weapons []domain.We `json:"weapons"`
	Skills       []domain.Skill              `json:"skills"`
}
