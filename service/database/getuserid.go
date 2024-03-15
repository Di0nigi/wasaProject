package database

func (db *appdbimpl) GetUserID(id string) (bool, error) {
	var result string
	err := db.c.QueryRow("SELECT UserId FROM Users WHERE UserId=?",id).Scan(&result)
	if err != nil {
		return false, err
	}
	return true, err
	
	
}