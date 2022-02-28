package db

import (
	"database/sql"
	"log"
)

//Represents user
type User struct {
	User_ID  int    `json:"id"`
	Username string `json:"username"`
}

//Represent uchat
type Chat struct {
	Chat_ID int    `json:"id"`
	Name    string `json:"name"`
	Users   int    `json:"users"`
}

//Query implementation that adds a user and returns id
func (u *User) addUser(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO \"user\" (username) VALUES($1) RETURNING user_id", u.Username).Scan(&u.User_ID)
	if err != nil {
		return err
	}
	return err
}

//Query implementation that updates an existing user and returns id
func (u *User) updateUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE \"user\" SET username=$1 WHERE user_id=$2 RETURNING user_id", u.Username, u.User_ID)
	if err != nil {
		return err
	}
	return err
}

//Query implementation that adds a chat and returns id
func (c *Chat) addChat(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO chat (name, users) VALUES($1, $2) RETURNING chat_id", c.Name, c.Users).Scan(&c.Chat_ID)
	if err != nil {
		return err
	}
	return err
}

//Query implementation that returns a chat of a specific user
func (c *Chat) getChat(db *sql.DB) error {
	var err error
	chat := db.QueryRow("SELECT chat_id, name, users FROM chat WHERE chat_id=$1", c.Chat_ID).Scan(&c.Chat_ID, &c.Name, &c.Users)
	if err != nil {
		log.Fatal("Problem selecting id, text from message  ", err.Error())
	}
	return chat
}

//Query implementation that delets a chat
func (c *Chat) deleteChat(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM chat WHERE chat_id=$1", c.Chat_ID)
	return err
}
