package database
import (
	"database/sql")

func (db *appdbimpl) GetAllphotos(id string) (error, []PostedImage) {
	var stream []PostedImage
	var rows *sql.Rows
	var err error

	rows, err = db.c.Query("SELECT * FROM PostedImages WHERE ownerId=?",id)
	
	for rows.Next() {
		var currentPost PostedImage
		err = rows.Scan(&currentPost.IDPhoto,&currentPost.Owner.ID,&currentPost.Image,&currentPost.Likes,&currentPost.NumComments)
		if err != nil {
        	return err, stream
    	}
		var commArray []Comment
		var commentsRows *sql.Rows
		
		commentsRows, err = db.c.Query("SELECT * FROM Comments WHERE PhotoId=?",currentPost.IDPhoto.IDObj)
		if err != nil {
        	return err, stream
    	}

		for commentsRows.Next(){
			var currentComment Comment
			commentsRows.Scan(&currentComment.IDComment.IDObj,&currentComment.Content,&currentComment.Owner.ID,&currentComment.Photo.IDObj)
			commArray=append(commArray,currentComment)
		}
		currentPost.Comments=commArray

		stream=append(stream,currentPost)

	}





	return err,stream



}