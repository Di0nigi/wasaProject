package database


func (db *appdbimpl) GetUserProfile(id string) (error, User) {
	var user User

	var userid UserId
	userid.IDUser=id
	user.ID=userid

	var followers []UserId
	rows, err := db.c.Query("SELECT followedBy FROM Followers WHERE UserId=?",id)
	if err!=nil{
		return err, user
	}

	for rows.Next(){
		var follower UserId
		rows.Scan(&follower)
		followers=append(followers,follower)
	}
	user.Followers=followers
	return err, user
}
