package database

func (db *appdbimpl) GetBanned(id string, bnId string) (error, UserId) {
	var user UserId

	err := db.c.QueryRow("SELECT UserId FROM BlockedUsers WHERE UserId=? AND BlockedBy=?", bnId, id).Scan(&user.IDUser)
	if err != nil {
		return err, user
	}
	return err, user
}
