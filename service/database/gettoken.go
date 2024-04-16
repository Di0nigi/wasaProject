package database

func (db *appdbimpl) GetToken(tk string) (error, string) {
	var result string
	err := db.c.QueryRow("SELECT UserId FROM Tokens WHERE token=?", tk).Scan(&result)
	if err != nil {

		return err, result
	}
	return err, result

}

func (db *appdbimpl) GetUserToken(user string) (error, string) {
	var result string
	err := db.c.QueryRow("SELECT token FROM Tokens WHERE UserId=?", user).Scan(&result)
	if err != nil {
		return err, result
	}
	return err, result

}
