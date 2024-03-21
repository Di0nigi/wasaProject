package database


func (db *appdbimpl) DelBlocked(blocked string , blockedby string ) (error) {
	ret, err := db.c.Exec("DELETE FROM BlockedUsers WHERE UserId = ? AND BlockedBy = ?", blocked, blockedby)

	
	return err
}