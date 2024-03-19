package database

import "fmt"

func (db *appdbimpl) DelFollower(follow string , followedby string ) (error) {
	fmt.Printf(follow)
	res, err := db.c.Exec("DELETE FROM Followers WHERE UserId = ? AND Followedby = ?", follow, followedby)

	rowsAffected, errs := res.RowsAffected()
    if errs != nil {
        return err
    }
	if (rowsAffected==0){
		fmt.Printf(follow+"|")
		fmt.Printf(" "+followedby)

	}
	//fmt.Printf(err)
	return err
}