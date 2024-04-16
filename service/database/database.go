/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserID(string) (bool, error)
	CreateUser(string) (error, string)
	ChangeUsername(string, string) error
	GetAllphotos(string) (error, []PostedImage)
	GetUserProfile(string) (error, User)
	AddBlocked(string, string) error
	DelBlocked(string, string) error
	AddFollower(string, string) error
	DelFollower(string, string) error
	AddPost(string, PostedImage) error
	DelPost(string, string) error
	AddComment(string, Comment) error
	DelComment(string, string, string) error
	AddLike(string, Like) error
	DelLike(string, string, string) error
	GetPhoto(string, string) (error, PostedImage)
	GetLike(string, string) (error, Like)
	GetFollower(string, string) (error, UserId)
	GetBanned(string, string) (error, UserId)
	GetToken(string) (error, string)
	GetUserToken(string) (error, string)
	/*GetComment(id api.ObjId)(api.Comment, error)
	GetBlocked(id api.UserId)(api.User, error)
	GetImage(id api.ObjId)(api.PostedImage, error)
	GetFollower(id api.UserId)(api.User, error)
	GetFollow(id api.UserId)(api.User, error)
	GetImages(id api.UserId)


	CreateUserUser(user api.User) error
	SaveComment(comment api.Comment) error
	SaveBlocked(blocked api.User,user api.User) error
	SaveImage(img api.PostedImage) error
	SaveFollower(user api.User,followedby api.User) error*/

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {

	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	_, err := db.Exec("PRAGMA foreign_key = ON")
	if err != nil {
		return nil, err
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	query := `CREATE TABLE IF NOT EXISTS Users (UserId TEXT PRIMARY KEY);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	query = `CREATE TABLE IF NOT EXISTS Tokens (token TEXT PRIMARY KEY, UserId TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	query = `INSERT INTO Tokens (token, UserId) SELECT ?, ? WHERE NOT EXISTS (
		SELECT 1 FROM Tokens WHERE token = ? AND UserId = ? );`
	_, err = db.Exec(query, "", "", "", "")
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	query = `CREATE TABLE IF NOT EXISTS BlockedUsers (tableId INTEGER PRIMARY KEY AUTOINCREMENT, UserId TEXT, BlockedBy TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	query = `CREATE TABLE IF NOT EXISTS Followers (tableId INTEGER PRIMARY KEY AUTOINCREMENT, UserId TEXT, Followedby TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	query = `CREATE TABLE IF NOT EXISTS PostedImages (PhotoId TEXT PRIMARY KEY, ownerId TEXT,imageData VARBINARY(100000), likes INTEGER, numComments INTEGER, date TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	query = `CREATE TABLE IF NOT EXISTS Comments (CommentId TEXT PRIMARY KEY, ownerId TEXT, content TEXT, PhotoId TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	query = `CREATE TABLE IF NOT EXISTS Likes (LikeId TEXT PRIMARY KEY, ownerId TEXT, PhotoId TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
