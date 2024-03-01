package proficiency

var (
	QueryInsertProficiency = `INSERT INTO proficiency (name, type) VALUES(?,?);`
	QueryGetAll            = `SELECT * from proficiency;`
	QueryGetById           = `SELECT * FROM proficiency WHERE proficiency_id = ?;`
	QueryUpdate      = `UPDATE proficiency SET name = ?, type = ?;`
	QueryDelete       = `DELETE FROM proficiency WHERE proficiency_id = ?;`
)
