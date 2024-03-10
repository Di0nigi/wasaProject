package database


func (db *appdbimpl) AddBlocked(blocked string , blockedby string ) (error) {
	_, err := db.c.Exec("INSERT INTO BlockedUsers (UserId, BlockedBy) VALUES (?, ?)", blocked, blockedby)
	return err
}