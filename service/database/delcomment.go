package database


func (db *appdbimpl) DelComment(user string , comm Comment ) (error) {
	_, err := db.c.Exec("DELETE FROM Comments WHERE CommentId = ? AND PhotoId = ?", comm.IDComment, user)
	return err
}