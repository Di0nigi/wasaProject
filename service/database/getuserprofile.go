package database

//import "fmt"


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
		var us string
		var follower UserId
		rows.Scan(&us)
		follower.IDUser=us
		followers=append(followers,follower)
	}
	user.Followers=followers
	

	var follows []UserId
	rows, err = db.c.Query("SELECT UserId FROM Followers WHERE FollowedBy=?",id)
	if err!=nil{
		return err, user
	}
	for rows.Next(){
		var followed UserId
		var use string
		rows.Scan(&use)
		followed.IDUser=use
		follows=append(follows,followed)
	}
	user.Follows=follows

	var blocked []UserId
	rows, err = db.c.Query("SELECT UserId FROM BlockedUsers WHERE BlockedBy=?",id)
	for rows.Next(){
		var blockeduser string
		var blkUser  UserId

		rows.Scan(&blockeduser)
		blkUser.IDUser=blockeduser
		blocked=append(blocked,blkUser)
	}
	user.BlockedArray=blocked

	var posts []ObjId
	rows, err = db.c.Query("SELECT PhotoId FROM PostedImages WHERE ownerId=?",id)
	for rows.Next(){
		var post ObjId
		rows.Scan(&post)
		posts=append(posts,post)
	}
	user.Posts=posts

	return err, user
}
