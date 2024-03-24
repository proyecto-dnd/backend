package user

var (
	QueryInsertUser = `INSERT INTO user (uid, name, email, password, display_name) VALUES(?,?,?,?,?)`
	//user.Id, &user.Username, &user.Email, &user.Password, &user.DisplayName
	QueryGetAllUsers = `SELECT uid, name, email, display_name, image FROM user`
	QueryGetUserById = `SELECT * FROM user WHERE user_id = ?`
	QueryUpdateUser  = `UPDATE user SET username = ?, email = ?, password = ? WHERE user_id = ?`
	QueryDeleteUser  = `DELETE FROM user WHERE user_id = ?`
)
