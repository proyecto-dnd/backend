package character

var (
	QueryInsertCharacter = `INSERT INTO table_character(name, class, race, background, hitpoints, 
	speed, armor_class, level, exp, str, dex, con, int, wiz, cha) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	QueryGetAllCharacters = `SELECT * FROM table_character`
	QueryGetCharacterById = `SELECT * FROM table_character WHERE character_id = ?`
	QueryUpdateCharacter  = `UPDATE table_character SET name = ?, class = ?, race = ?, background = ?, 
	hitpoints = ?, speed = ?, armor_class = ?, level = ?, exp = ?, str = ?, dex  ?, con = ?, int = ?, wiz = ?, cha = ? WHERE character_id = ?`
	QueryDeleteCharacter = `DELETE FROM table_character WHERE character_id = ?`
)
