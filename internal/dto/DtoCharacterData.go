package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type FullCharacterData struct {
	CharacterId int64  `json:"idcharacter"`
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	//Creo que conviene hacer una tabla de clases para tener dropdowns en el front
	Class int `json:"class"`
	// Lo mismo que con class
	Race       int                         `json:"race"`
	Background string                      `json:"background"`
	Hitpoints  int                         `json:"hitpoints"`
	Speed      int                         `json:"speed"`
	ArmorClass int                         `json:"armor_class"`
	Level      int                         `json:"level"`
	Exp        int                         `json:"exp"`
	CampaignId int                         `json:"campaign_id"`
	Str        int                         `json:"str"`
	Dex        int                         `json:"dex"`
	Int        int                         `json:"int"`
	Wiz        int                         `json:"wiz"`
	Con        int                         `json:"con"`
	Cha        int                         `json:"cha"`
	Items      []domain.ItemXCharacterData `json:"items"`
	Skills     []domain.Skill              `json:"skills"`
}