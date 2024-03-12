package database
import (
	"database/sql";
"fmt"
)

func (db *appdbimpl) GetAllphotos(id string) (error, []PostedImage) {
	var stream []PostedImage
	var rows *sql.Rows
	var err error

	rows, err = db.c.Query("SELECT * FROM PostedImages WHERE ownerId=?",id)
	
	for rows.Next() {
		var currentPost PostedImage
		err = rows.Scan(&currentPost.IDPhoto.IDObj,&currentPost.Owner.IDUser,&currentPost.Image,&currentPost.Likes,&currentPost.NumComments)
		if err != nil {
        	return err, stream
    	}
		var commArray []Comment
		//var commentsRows *sql.Rows
		
		commentsRows, errs := db.c.Query("SELECT * FROM Comments WHERE ownerId = ?",id)
		if errs != nil {
        	return err, stream
    	}
		

		for commentsRows.Next(){
			var currentComment Comment
			commentsRows.Scan(&currentComment.IDComment.IDObj,&currentComment.Content,&currentComment.Owner.IDUser,&currentComment.Photo.IDObj)
			fmt.Printf(currentComment.Content)
			commArray=append(commArray,currentComment)
		}
		currentPost.Comments=commArray
		//fmt.Printf(err)
		//fmt.Printf(commArray[0].Content)

		stream=append(stream,currentPost)

	}





	return err,stream



}