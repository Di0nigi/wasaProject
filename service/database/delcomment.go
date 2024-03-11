package database


func (db *appdbimpl) DelComment(user string , comm Comment ) (error) {
	_, err := db.c.Exec("DELETE FROM Comments WHERE CommentId = ? AND PhotoId = ?", comm.IDComment.IDObj, comm.Photo.IDObj)
	if err!=nil{
		return err
	}

	var nComm int
	err = db.c.QueryRow("SELECT numComments FROM PostedImages WHERE PhotoId = ?", comm.Photo.IDObj).Scan(&nComm)
	if err!=nil{
		return err
	}
	nComm-=1
	_, err = db.c.Exec("UPDATE PostedImages SET numComments = ? WHERE PhotoId = ?", nComm, comm.Photo.IDObj)

	
	return err
}