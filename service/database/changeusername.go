package database

func (db *appdbimpl) ChangeUsername(oldId string, newId string) error {

	_, err := db.c.Exec(`UPDATE Users SET UserId=? WHERE UserId=?`, newId, oldId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE BlockedUsers SET UserId=? WHERE UserId=?`, newId, oldId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE BlockedUsers SET BlockedBy=? WHERE BlockedBy=?`, newId, oldId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`UPDATE Followers SET UserId=? WHERE UserId=?`, newId, oldId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE Followers SET FollowedBy=? WHERE FollowedBy=?`, newId, oldId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`UPDATE PostedImages SET ownerId=? WHERE ownerId=?`, newId, oldId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE Comments SET ownerId=? WHERE ownerId=?`, newId, oldId)
	if err != nil {
		return err
	}

	return err
}
