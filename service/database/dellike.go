package database


func (db *appdbimpl) DelLike(user string , lk Like ) (error) {
	_, err := db.c.Exec("DELETE FROM Likes WHERE LikeId = ? AND PhotoId = ?", lk.IDLike.IDObj, lk.ToPhoto.IDObj)
	if err!=nil{
		return err
	}

	var nLike int
	err = db.c.QueryRow("SELECT likes FROM PostedImages WHERE PhotoId = ?", lk.ToPhoto.IDObj).Scan(&nLike)
	if err!=nil{
		return err
	}
	nLike-=1
	_, err = db.c.Exec("UPDATE PostedImages SET likes = ? WHERE PhotoId = ?", nLike, lk.ToPhoto.IDObj)

	return err
}