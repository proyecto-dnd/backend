package characterXspell

var (
	QueryInsert = `INSERT INTO character_spell (character_id, spell_id) values(?,?);`
	QueryDelete = `DELETE FROM character_spell WHERE character_spell_id=?;`
)
