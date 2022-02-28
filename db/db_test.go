package db

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var s Service

func TestMain(m *testing.M) {
	s = Service{}

	//load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading env variables")
	}

	s.Init(
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_NAME"),
		os.Getenv("PG_SSL"))

	code := m.Run() //executes tests

	os.Exit(code)
}

//Testing response to the /user/add endpoint
func TestAddUser(t *testing.T) {
	var jsonStr = []byte(`{"username": "John"}`)

	//sending POST request to the endpoint
	req, err := http.NewRequest("POST", "/user/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//making sure we get expected status code
	r := httptest.NewRecorder()
	s.Router.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	//checking body
	expected := `{"id":1,"username":"John"}` //expected response example
	if r.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", r.Body.String(), expected)
	}
}

//Testing response to the /chat/add endpoint
func TestAddChat(t *testing.T) {
	var jsonStr = []byte(`{"name": "chat_1", "users": 1}`)

	//sending POST request to the endpoint
	req, err := http.NewRequest("POST", "/chat/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//making sure we get expected status code
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	expected := `{"id":1,"name":"chat_1","users":1}` //expected response example
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

//Testing response to the /user/update/1 endpoint
func TestUpdateUser(t *testing.T) {
	var jsonStr = []byte(`{"username": "Bob"}`)

	//sending put request to the endpoint
	req, err := http.NewRequest("PUT", "/user/update/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//making sure we get expected status code
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	expected := `{"id":1,"username":"Bob"}` //expected response example
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

//Testing response to the /chat/get/1 endpoint
func TestGetChat(t *testing.T) {
	//sending GET request to the endpoint
	req, err := http.NewRequest("GET", "/chat/get/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	//making sure we get expected status code
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"id":1,"name":"chat_1","users":1}` //expected response example
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

//Testing response to the /chat/delete/1 endpoint
func TestDeleteChat(t *testing.T) {
	//sending DELETE request to the endpoint
	req, err := http.NewRequest("DELETE", "/chat/delete/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	//making sure we get expected status code
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
	expected := `"successfully deleted chat!"` //expected response example
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v expected %v", rr.Body.String(), expected)
	}
}
