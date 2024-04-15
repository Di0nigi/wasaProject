package database

func (db *appdbimpl) AddPost(owner string, image PostedImage) error {

	_, err := db.c.Exec("INSERT INTO PostedImages (PhotoId, ownerId, imageData, likes, numComments, date) VALUES (?, ?, ?, ?, ?, ?)", image.IDPhoto.IDObj, owner, image.Image, image.Likes, image.NumComments, image.Date)

	return err
}
