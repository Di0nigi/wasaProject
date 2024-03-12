package database


func (db *appdbimpl) DelLike(user string , lk string, photo string) (error) {
	_, err := db.c.Exec("DELETE FROM Likes WHERE LikeId = ? AND PhotoId = ?", lk, photo)
	if err!=nil{
		return err
	}

	var nLike int
	err = db.c.QueryRow("SELECT likes FROM PostedImages WHERE PhotoId = ?", photo).Scan(&nLike)
	if err!=nil{
		return err
	}
	nLike-=1
	_, err = db.c.Exec("UPDATE PostedImages SET likes = ? WHERE PhotoId = ?", nLike, photo)

	return err
}