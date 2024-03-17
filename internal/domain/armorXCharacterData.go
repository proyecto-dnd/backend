package domain

type ArmorXCharacterData struct {
	ArmorXCharacterData_Id int   `json:"armorxcharacterdataid"`
	Armor               Armor `json:"armor"`
	CharacterData_Id       int   `json:"characterdataid"`
	Equipped               bool  `json:"equipped"`
}
