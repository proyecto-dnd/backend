package user

var (
	QueryInsertUser = `INSERT INTO user (uid, name, email, password, display_name,image) VALUES(?,?,?,?,?,?)`	
	QueryGetAllUsers             = `SELECT uid, name, email, display_name, image FROM user`
	QueryGetUserById             = `SELECT uid, name, email, display_name, image FROM user WHERE uid = ?`
	QueryUpdateUser              = `UPDATE user SET name = ?, email = ?, password = ?, image= ?, display_name = ? WHERE user_id = ?`
	QueryDeleteUser              = `DELETE FROM user WHERE user_id = ?`
	QueryGetSubExpirationDate    = `SELECT sub_expiration FROM user WHERE uid = ?`
	QueryUpdateSubExpirationDate = `UPDATE user SET sub_expiration = ? WHERE uid = ?`
)
