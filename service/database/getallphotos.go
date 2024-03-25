package database


func (db *appdbimpl) GetAllphotos(id string) (error, []PostedImage) {
	var stream []PostedImage
	// var rows *sql.Rows
	// var err error

	rows, err := db.c.Query("SELECT * FROM PostedImages WHERE ownerId=?",id)

	if err != nil {
		return err, stream
	}
	
	for rows.Next() {
		var currentPost PostedImage
		err = rows.Scan(&currentPost.IDPhoto.IDObj,&currentPost.Owner.IDUser,&currentPost.Image,&currentPost.Likes,&currentPost.NumComments)
		if err != nil {
        	return err, stream
    	}
		var commArray []Comment
	
		
		commentsRows, errs := db.c.Query("SELECT * FROM Comments WHERE ownerId = ?",id)
		if errs != nil {
        	return err, stream
    	}
		
		for commentsRows.Next(){
			var currentComment Comment
			err =commentsRows.Scan(&currentComment.IDComment.IDObj,&currentComment.Owner.IDUser,&currentComment.Content,&currentComment.Photo.IDObj)
			if err != nil {
				return err, stream
			}
			
			commArray=append(commArray,currentComment)
		}
		if err = commentsRows.Err(); err != nil {
			return err, stream
		}
		currentPost.Comments=commArray
		
		stream=append(stream,currentPost)
	}
	if err = rows.Err(); err != nil {
		return err, stream
	}





	return err,stream



}