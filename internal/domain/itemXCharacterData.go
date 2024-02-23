package domain

type ItemXCharacterData struct {
	ItemXCharacterData_Id int64 `json:"itemxcharacterdataid"`
	Item Item `json:"item"`
	CharacterData_Id int64 `json:"characterdataid"`
	Quantity int `json:"quantity"`
}