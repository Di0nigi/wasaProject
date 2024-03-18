package database
import (
	"database/sql"
	"fmt"

)

func (db *appdbimpl) GetPhoto(id string, photoId string) (error, PostedImage) {

	var rows *sql.Rows
	var err error
	var currentPost PostedImage

	rows, err = db.c.Query("SELECT * FROM PostedImages WHERE ownerId=? AND PhotoId=?",id, photoId)
	if err != nil {
		return err, currentPost
	}
	for rows.Next() {
	err = rows.Scan(&currentPost.IDPhoto.IDObj,&currentPost.Owner.IDUser,&currentPost.Image,&currentPost.Likes,&currentPost.NumComments)
		if err != nil {
        	return err, currentPost
    	}
		var commArray []Comment
		//var commentsRows *sql.Rows
		
		commentsRows, errs := db.c.Query("SELECT * FROM Comments WHERE ownerId = ? AND PhotoId=?",id,photoId)
		if errs != nil {
        	return err, currentPost
    	}
		
		for commentsRows.Next(){
			var currentComment Comment
			err =commentsRows.Scan(&currentComment.IDComment.IDObj,&currentComment.Owner.IDUser,&currentComment.Content,&currentComment.Photo.IDObj)
			fmt.Printf(currentComment.Content)
			fmt.Printf(currentComment.Owner.IDUser)
			if err != nil {
				return err, currentPost
			}
			
			commArray=append(commArray,currentComment)
		}
		if err = commentsRows.Err(); err != nil {
			return err, currentPost
		}
		currentPost.Comments=commArray
	}

	return err, currentPost

}
	