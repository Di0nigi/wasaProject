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
		err=rows.Scan(&us)
		if err!=nil{
			return err, user
		}
		follower.IDUser=us
		followers=append(followers,follower)
	}
	if err = rows.Err(); err != nil {
		return err, user
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
		err=rows.Scan(&use)
		if err!=nil{
			return err, user
		}
		followed.IDUser=use
		follows=append(follows,followed)
	}
	if err = rows.Err(); err != nil {
		return err, user
	}
	user.Follows=follows

	var blocked []UserId
	rows, err = db.c.Query("SELECT UserId FROM BlockedUsers WHERE BlockedBy=?",id)
	for rows.Next(){
		var blockeduser string
		var blkUser  UserId

		err=rows.Scan(&blockeduser)
		if err!=nil{
			return err, user
		}
		blkUser.IDUser=blockeduser
		blocked=append(blocked,blkUser)
	}
	if err = rows.Err(); err != nil {
		return err, user
	}
	user.BlockedArray=blocked
	

	var posts []ObjId
	rows, err = db.c.Query("SELECT PhotoId FROM PostedImages WHERE ownerId=?",id)
	for rows.Next(){
		var post ObjId
		var stPost string
		err =rows.Scan(&stPost)
		if err!=nil{
			return err, user
		}
		post.IDObj=stPost
		posts=append(posts,post)
	}
	if err = rows.Err(); err != nil {
		return err, user
	}
	user.Posts=posts
	

	return err, user
}
