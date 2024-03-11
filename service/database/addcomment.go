package database


func (db *appdbimpl) AddComment(user string, comm Comment ) (error) {
	_, err := db.c.Exec("INSERT INTO Comments (CommentId, ownerId, content, PhotoId) VALUES (?, ?, ?, ?)",comm.IDComment, user, comm.Content, user, comm.Photo)
	if err!=nil{
		return err
	}
	
	var nComm int
	err = db.c.QueryRow("SELECT numComments FROM PostedImages WHERE PhotoId = ?", comm.Photo).Scan(&nComm)
	if err!=nil{
		return err
	}
	nComm+=1
	_, err = db.c.Exec("UPDATE PostedImages SET numComments = ? WHERE PhotoId = ?", nComm, comm.Photo)

	
	return err
}