package database

type UserID struct {
	IDUser string `json:"idUser" minLength:"3" maxLength:"16" pattern:"^.*?$"`
}