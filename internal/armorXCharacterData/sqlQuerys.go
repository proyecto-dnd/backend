package armorXCharacterData

var (
	QueryCreateCharacterArmor              = `INSERT INTO character_armor (character_id, armor_id, equipped) VALUES (?, ?, ?)`
	QueryGetAllCharacterArmor              = `SELECT * FROM character_armor`
	QueryGetByIdCharacterArmor             = `SELECT * FROM character_armor WHERE character_armor_id = ?`
	QueryGetByCharacterIdCharacterArmor    = `SELECT * FROM character_armor WHERE character_id = ?`
	QueryUpdateCharacterArmor              = `UPDATE character_armor SET character_id = ?, armor_id = ?, equipped = ? WHERE character_armor_id = ?`
	QueryDeleteCharacterArmor              = `DELETE FROM character_armor WHERE character_armor_id = ?`
	QueryDeleteByCharacterIdCharacterArmor = `DELETE FROM character_armor WHERE character_id = ?`
)
