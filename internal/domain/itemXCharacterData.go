package domain

type ItemXCharacterData struct {
	ItemXCharacterDataId int64 `json:"itemxcharacterdataid"`
	Item Item `json:"item"`
	CharacterDataId int64 `json:"characterdataid"`
	Quantity int `json:"quantity"`
}