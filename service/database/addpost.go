package database


func (db *appdbimpl) AddPost(owner string , image PostedImage) (error) {
	_, err := db.c.Exec("INSERT INTO PostedImages (PhotoId, ownerId, imageData, likes, numComments) VALUES (?, ?, ?, ?, ?)", image.IDPhoto, owner, image.Image, image.Likes, image.NumComments)
	return err
}