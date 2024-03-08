package database

func (db *appdbimpl) CreateUser(id string) (error) {
	_, err := db.c.Exec("INSERT INTO Users (UserId) VALUES (?)", id)
	return err
}
