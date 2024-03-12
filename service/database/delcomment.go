package database


func (db *appdbimpl) DelComment(user string , comm string, photo string ) (error) {
	_, err := db.c.Exec("DELETE FROM Comments WHERE CommentId = ? AND PhotoId = ?", comm, photo)
	if err!=nil{
		return err
	}

	var nComm int
	err = db.c.QueryRow("SELECT numComments FROM PostedImages WHERE PhotoId = ?", photo).Scan(&nComm)
	if err!=nil{
		return err
	}
	nComm-=1
	_, err = db.c.Exec("UPDATE PostedImages SET numComments = ? WHERE PhotoId = ?", nComm, photo)

	
	return err
}