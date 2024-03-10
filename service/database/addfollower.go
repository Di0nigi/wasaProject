package database


func (db *appdbimpl) AddFollower(tofollow string , followedby string ) (error) {
	_, err := db.c.Exec("INSERT INTO Followers (UserId, FollowedBy) VALUES (?, ?)", tofollow, followedby)
	return err
}