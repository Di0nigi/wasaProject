package database


func (db *appdbimpl) AddComment(user string, comm Comment ) (error) {
	_, err := db.c.Exec("INSERT INTO Comments (CommentId, ownerId, content, PhotoId) VALUES (?, ?, ?, ?)",comm.IDComment, user, comm.Content, user, comm.Photo)
	return err
}