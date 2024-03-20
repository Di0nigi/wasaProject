package database


func (db *appdbimpl) DelFollower(follow string , followedby string ) (error) {
	
	res, err := db.c.Exec("DELETE FROM Followers WHERE UserId = ? AND Followedby = ?", follow, followedby)

	rowsAffected, errs := res.RowsAffected()
    if errs != nil {
        return err
    }
	if (rowsAffected==0){
		

	}
	
	return err
}