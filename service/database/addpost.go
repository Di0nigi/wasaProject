package database
import "fmt"


func (db *appdbimpl) AddPost(owner string , image PostedImage) (error) {
	fmt.Printf(image.Image)
	fmt.Printf(" reached")
	_, err := db.c.Exec("INSERT INTO PostedImages (PhotoId, ownerId, imageData, likes, numComments) VALUES (?, ?, ?, ?, ?)", image.IDPhoto.IDObj, owner, image.Image, image.Likes, image.NumComments)
	return err
}