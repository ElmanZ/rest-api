package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Service struct {
	Router *mux.Router
	DB     *sql.DB
}

type User struct {
	User_ID   int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
	fmt.Println("Go")

	s := Service{}
	s.init()
	s.Start("HTTP_PORT")
}
func (s *Service) Routes() {
	s.Router.HandleFunc("/users/add", s.addUser).Methods("POST")
}

func (u *User) addUser(db *sql.DB) error {
	return db.QueryRow("INSER INTO 'user' (username) VALUES ($1) RETURNING user_id", u.Username).Scan(&u.User_ID)
}

func (s *Service) addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "application/json")
	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		log.Fatal(err, "Invalid request")
	}
	defer r.Body.Close()
	if err := u.addUser(s.DB); err != nil {
		return
	}
	fmt.Fprintf(w, "User with id %d, created successfully\n", u.User_ID)
}

func (s *Service) Start(addr string) {
	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_PORT"), s.Router))
}

func (s *Service) init() {
	var pg = os.Getenv("PG_URL")
	var err error
	s.DB, err = sql.Open("postgres", pg)
	if err != nil {
		log.Fatal(err)
	}

	defer s.DB.Close()

	s.Router = mux.NewRouter()
	s.init()
}
