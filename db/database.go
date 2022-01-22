package db

import (
	"database/sql"
	"time"
)

//Represents user
type User struct {
	User_ID   int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

//Query implementation that adds a user and returns id
func (u *User) addUser(db *sql.DB) error {
	return db.QueryRow("INSERT INTO \"user\" (username) VALUES ($1) RETURNING user_id", u.Username).Scan(&u.User_ID)
}
