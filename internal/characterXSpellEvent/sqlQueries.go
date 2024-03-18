package characterxspellevent

var (
	QueryGetAll = `SELECT * FROM character_spell_event;`
	QueryGetById = `SELECT * FROM character_spell_event WHERE character_spell_event_id=?;`
	QueryGetByCharacterId = `SELECT * FROM character_spell_event WHERE character_id=?;`
	QueryGetBySpellEventId = `SELECT * FROM character_spell_event WHERE event_id=?;`
	QueryInsert = `INSERT INTO character_spell (character_id, spell_id) values(?,?);`
	QueryDelete = `DELETE FROM character_spell WHERE character_spell_id=?;`
)