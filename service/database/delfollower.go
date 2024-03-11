package database

import "fmt"

func (db *appdbimpl) DelFollower(follow string , followedby string ) (error) {
	fmt.Printf(follow)
	_, err := db.c.Exec("DELETE FROM Followers WHERE UserId = ? AND Followedby = ?", follow, followedby)
	//fmt.Printf(err)
	return err
}