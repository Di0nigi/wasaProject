package database


func (db *appdbimpl) AddFollower(tofollow string , followedby string ) (error) {
	_, err := db.c.Exec("INSERT INTO Followers (UserId, Followedby) VALUES (?, ?)", tofollow, followedby)
	return err
}