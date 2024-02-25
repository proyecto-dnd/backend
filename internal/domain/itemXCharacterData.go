package domain

type ItemXCharacterData struct {
	ItemXCharacterData_Id int64 `json:"itemxcharacterdataid"`
	CharacterData_Id int64 `json:"characterdataid"`
	Item Item `json:"item"`
	Quantity int `json:"quantity"`
}