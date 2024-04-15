package database

import "fmt"

func (db *appdbimpl) GetPhoto(id string, photoId string) (error, PostedImage) {

	var currentPost PostedImage

	rows, err := db.c.Query("SELECT * FROM PostedImages WHERE ownerId=? AND PhotoId=?", id, photoId)
	if err != nil {
		return err, currentPost
	}
	for rows.Next() {
		err = rows.Scan(&currentPost.IDPhoto.IDObj, &currentPost.Owner.IDUser, &currentPost.Image, &currentPost.Likes, &currentPost.NumComments, &currentPost.date)
		fmt.Printf(currentPost.date)
		fmt.Printf("here is date")
		if err != nil {
			return err, currentPost
		}
		var commArray []Comment

		commentsRows, errs := db.c.Query("SELECT * FROM Comments WHERE PhotoId=?", photoId)
		if errs != nil {
			return err, currentPost
		}

		for commentsRows.Next() {
			var currentComment Comment
			err = commentsRows.Scan(&currentComment.IDComment.IDObj, &currentComment.Owner.IDUser, &currentComment.Content, &currentComment.Photo.IDObj)
			if err != nil {
				return err, currentPost
			}

			commArray = append(commArray, currentComment)
		}
		if err = commentsRows.Err(); err != nil {
			return err, currentPost
		}
		currentPost.Comments = commArray
	}
	if err = rows.Err(); err != nil {
		return err, currentPost

	}

	return err, currentPost

}
