package database

func (db *appdbimpl) GetUserID(id string) (bool, error) {
	var result bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT UserId FROM Users WHERE UserId = ?)", id).Scan(&result)
	if err != nil {
		result = false
		return result, err
	}
	return result, err

}
