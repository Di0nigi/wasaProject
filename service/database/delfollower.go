package database


func (db *appdbimpl) DelFollower(follow string , followedby string ) (error) {
	_, err := db.c.Exec("DELETE FROM Followers WHERE UserId = ? AND FollowedBy = ?", follow, followedby)
	return err
}