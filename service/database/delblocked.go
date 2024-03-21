package database

import "fmt"

func (db *appdbimpl) DelBlocked(blocked string , blockedby string ) (error) {
	ret, err := db.c.Exec("DELETE FROM BlockedUsers WHERE UserId = ? AND BlockedBy = ?", blocked, blockedby)

	rowsAffected, errs := ret.RowsAffected()
    if errs != nil {
        return err
    }
	if (rowsAffected==0){
		fmt.Printf("zero")
		

	}
	
	return err
}