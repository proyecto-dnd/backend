package characterxattackevent

var (
	QueryGetAll = `SELECT * FROM character_attack_event;`
	QueryGetById = `SELECT * FROM character_attack_event WHERE character_attack_event_id=?;`
	QueryGetByCharacterId = `SELECT * FROM character_attack_event WHERE character_id=?;`
	QueryGetBySpellEventId = `SELECT * FROM character_attack_event WHERE event_id=?;`
	QueryInsert = `INSERT INTO character_attack_event (character_id, event_id) values(?,?);`
	QueryDelete = `DELETE FROM character_attack_event character_attack_event_id=?;`
)