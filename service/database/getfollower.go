package database


func (db *appdbimpl) GetFollower(id string, fId string) (error, UserId) {
	var user UserId

	err := db.c.QueryRow("SELECT UserId FROM Followers WHERE UserId=? AND Followedby=?",fId, id).Scan(&user.IDUser)
	if err != nil {
		return err, user
	}
	return err, user
}