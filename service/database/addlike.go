package database


func (db *appdbimpl) AddLike(owner string , lk Like) (error) {
	_, err := db.c.Exec("INSERT INTO Likes(LikeId, ownerId, PhotoId) VALUES (?, ?, ?)", lk.IDLike, owner, lk.ToPhoto)
	if err!=nil{
		return err
	}
	var nLike int
	err = db.c.QueryRow("SELECT likes FROM PostedImages WHERE PhotoId = ?", lk.ToPhoto).Scan(&nLike)
	if err!=nil{
		return err
	}
	nLike+=1
	_, err = db.c.Exec("UPDATE PostedImages SET likes = ? WHERE PhotoId = ?", nLike, lk.ToPhoto)

	return err



}