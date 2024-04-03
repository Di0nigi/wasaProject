package database

func (db *appdbimpl) SetName(name string) error {
	_, err := db.c.Exec("INSERT INTO example_table (id, name) VALUES (1, ?)", name)
	return err
}
