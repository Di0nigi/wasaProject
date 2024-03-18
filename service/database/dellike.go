package database

import "fmt"

func (db *appdbimpl) DelLike(user string , lk string, photo string) (error) {
	res, err := db.c.Exec("DELETE FROM Likes WHERE LikeId = ? AND PhotoId = ?", lk, photo)
	if err!=nil{
		return err
	}
	rowsAffected, errs := res.RowsAffected()
    if errs != nil {
        return err
    }
	if (rowsAffected==0){
		fmt.Printf("0 ")
		fmt.Printf(lk+" ")
		fmt.Printf(user+" ")
		fmt.Printf(photo+" ")
	}
	

	var nLike int
	err = db.c.QueryRow("SELECT likes FROM PostedImages WHERE PhotoId = ?", photo).Scan(&nLike)
	if err!=nil{
		return err
	}
	nLike-=1
	_, err = db.c.Exec("UPDATE PostedImages SET likes = ? WHERE PhotoId = ?", nLike, photo)

	return err
}