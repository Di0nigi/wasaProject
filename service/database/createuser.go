package database

import (
	"crypto/rand"
	"encoding/base64"
)

func (db *appdbimpl) CreateUser(id string) (error, string) {
	tk, err := generateRandomString(16)
	if err != nil {
		return err, tk
	}
	_, err = db.c.Exec("INSERT INTO Users (UserId) VALUES (?)", id)
	if err != nil {
		return err, tk
	}

	_, err = db.c.Exec("INSERT INTO Tokens (token, UserId) VALUES (?,?)", tk, id)
	return err, tk
}

func generateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b)[:n], nil
}
