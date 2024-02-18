package domain

type ItemXCharacterData struct {
	ItemXCharacterDataId int64 `json:"itemxcharacterdataid"`
	Item Item `json:"item"`
	CharacterDataId int64 `json:"characterdataid"`
	Equipped bool `json:"equipped"`
	Quantity int `json:"quantity"`
}