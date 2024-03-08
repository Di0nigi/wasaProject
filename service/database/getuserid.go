package database

func (db *appdbimpl) GetUserID(id string) (bool, error) {

	err := db.c.QueryRow("SELECT UserId FROM Users WHERE UserId=?").Scan(&id)
	if err != nil {
		return false, err
	}
	
	return true, err
	
	
}