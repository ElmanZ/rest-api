package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" // this package is used to handle url routes
	_ "github.com/lib/pq"    // pstgres driver
)

//Service struture allows us to expose router and database
type Service struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize connection whith postgres database
func (s *Service) Init(host, port, user, password, database, sslmode string) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=%s", host, port, user, password, database, sslmode)
	var err error
	s.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	s.Router = mux.NewRouter()
	s.routes()
}

//Defining routes that will use handlers
func (s *Service) routes() {
	s.Router.HandleFunc("/users/add", s.addUser).Methods("POST")
}

//This method start the application
func (s *Service) Start(addr string) {
	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_PORT"), s.Router))
}

//This method is a handler to add a user to the database and return the id
func (s *Service) addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "application/json")
	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		log.Fatal(err, "Invalid request")
	}

	if err := u.addUser(s.DB); err != nil {
		return
	}
	fmt.Fprintf(w, "User with id %d, created successfully\n", u.User_ID)
}

//Add other crud operations
