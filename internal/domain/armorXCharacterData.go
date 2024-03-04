package domain

type ArmorXCharacterData struct {
	ArmorXCharacterData_Id int64 `json:"armorxcharacterdataid"`
	Armor_Id               Armor `json:"armor"`
	CharacterData_Id       int64 `json:"characterdataid"`
	Equipped               bool  `json:"equipped"`
}
