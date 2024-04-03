package database

func (db *appdbimpl) DelPost(user string, photo string) error {
	_, err := db.c.Exec("DELETE FROM PostedImages WHERE PhotoId = ? AND ownerId = ?", photo, user)

	if err != nil {
		return err
	}
	_, err = db.c.Exec("DELETE FROM Comments WHERE PhotoId = ?", photo)

	return err
}
