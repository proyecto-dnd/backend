package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type NewTableCharacter struct {
	Idcharacter int    `json:"idcharacter"`
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	//Creo que conviene hacer una tabla de clases para tener dropdowns en el front
	Class string `json:"class"`
	// Lo mismo que con class
	Race        string     `json:"race"`
	Background  string     `json:"background"`
	Hitpoints   int        `json:"hitpoints"`
	Speed       int        `json:"speed"`
	Armor_class int        `json:"armor_class"`
	Level       int        `json:"level"`
	Exp         int        `json:"exp"`
	Campaign_id int        `json:"campaign_id"`
	Str         int        `json:"str"`
	Dex         int        `json:"dex"`
	Int         int        `json:"int"`
	Wiz         int        `json:"wiz"`
	Con         int        `json:"con"`
	Cha         int        `json:"cha"`
	Items       []domain.Item `json:"items"`
	Skills      []domain.Skill    `json:"skills"`
}