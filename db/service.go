package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux" // this package is used to handle url routes
	_ "github.com/lib/pq"    // postgres driver
)

//Service struture allows us to expose a router and a database
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
	s.Router.HandleFunc("/user/add", s.addUser).Methods("POST")
	s.Router.HandleFunc("/chat/add", s.addChat).Methods("POST")
	s.Router.HandleFunc("/user/update/{id}", s.UpdateUser).Methods("PUT")
	s.Router.HandleFunc("/chat/get/{id}", s.getChat).Methods("GET")
	s.Router.HandleFunc("/chat/delete/{id}", s.deleteChat).Methods("DELETE")
}

//This method start the application
func (s *Service) Start(addr string) {
	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_PORT"), s.Router))
}

//Method that will procces erros
func errResp(rw http.ResponseWriter, code int, message string) {
	jsonResp(rw, code, message)
}

//Method will process good responses
func jsonResp(rw http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)                               //marshaling the struct
	rw.Header().Set("Content-Type", "application/json; charset=utf-8") //setting up response headers for http req
	rw.WriteHeader(code)                                               //http status code
	rw.Write(response)                                                 //json response writer
}

//A handler to add a user to the database and return the id
func (s *Service) addUser(rw http.ResponseWriter, r *http.Request) {
	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		errResp(rw, http.StatusBadRequest, "Invalid request")
	}
	defer r.Body.Close()

	if err := u.addUser(s.DB); err != nil {
		errResp(rw, http.StatusInternalServerError, err.Error())
	}
	jsonResp(rw, http.StatusCreated, u)
}

//A handler to create a chat between users and return chat id
func (s *Service) addChat(rw http.ResponseWriter, r *http.Request) {
	var c Chat
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		errResp(rw, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	if err := c.addChat(s.DB); err != nil {
		errResp(rw, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResp(rw, http.StatusCreated, c)
}

//Updates existing user
func (s *Service) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		errResp(rw, http.StatusBadRequest, "Invalid User ID")
		log.Fatal(err.Error())
	}

	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		errResp(rw, http.StatusBadRequest, "Invalid request")
	}
	defer r.Body.Close()

	u.User_ID = id
	if err := u.updateUser(s.DB); err != nil {
		errResp(rw, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResp(rw, http.StatusCreated, u)
}

//A handler to get all chats of a user
func (s *Service) getChat(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		errResp(rw, http.StatusBadRequest, "Invalid Chat ID")
		log.Fatal(err.Error())
	}

	c := Chat{Chat_ID: (id)}
	if err := c.getChat(s.DB); err != nil {
		errResp(rw, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResp(rw, http.StatusOK, c)
}

//A handler to delete the chat from database
func (s *Service) deleteChat(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		errResp(rw, http.StatusBadRequest, "Invalid Chat ID")
		return
	}

	c := Chat{Chat_ID: (id)}
	if err := c.deleteChat(s.DB); err != nil {
		errResp(rw, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResp(rw, http.StatusOK, "successfully deleted chat!")
}
