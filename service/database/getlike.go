package database


func (db *appdbimpl) GetLike(id string, phId string) (error, Like) {
	var lk Like
	var idlk ObjId
	var idus UserId
	var phoId ObjId

	err := db.c.QueryRow("SELECT * FROM Likes WHERE ownerId=? AND PhotoId=?",id, phId).Scan(&idlk.IDObj,&idus.IDUser,&phoId.IDObj)
	if err != nil {
		return err, lk
	}
	lk.IDLike=idlk
	lk.Owner=idus
	lk.ToPhoto=phoId
	return err, lk
}
