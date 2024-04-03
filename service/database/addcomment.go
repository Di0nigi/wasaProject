package database

func (db *appdbimpl) AddComment(user string, comm Comment) error {

	_, err := db.c.Exec("INSERT INTO Comments (CommentId, ownerId, content, PhotoId) VALUES (?, ?, ?, ?)", comm.IDComment.IDObj, comm.Owner.IDUser, comm.Content, comm.Photo.IDObj)
	if err != nil {
		return err
	}

	var nComm int
	err = db.c.QueryRow("SELECT numComments FROM PostedImages WHERE PhotoId = ?", comm.Photo.IDObj).Scan(&nComm)
	if err != nil {
		return err
	}
	nComm += 1
	_, err = db.c.Exec("UPDATE PostedImages SET numComments = ? WHERE PhotoId = ?", nComm, comm.Photo.IDObj)

	return err
}
