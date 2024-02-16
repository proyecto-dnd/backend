package user

var (
	QueryInsertUser  = `INSERT INTO user(username, email, password) VALUES(?,?,?)`
	QueryGetAllUsers = `SELECT * FROM user`
	QueryGetUserById = `SELECT * FROM user WHERE user_id = ?`
	QueryUpdateUser  = `UPDATE user SET username = ?, email = ?, password = ? WHERE user_id = ?`
	QueryDeleteUser  = `DELETE FROM user WHERE user_id = ?`	
)
