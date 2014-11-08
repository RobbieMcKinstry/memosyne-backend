package main

import (
	"github.com/gorilla/mux"

	"fmt"
	"log"
	"net/http"
)

// TODO 
func HandleContactsRead(w http.ResponseWriter, r *http.Request)   {
	w.Write([]byte(`
		{ "contacts": [{"name": "Malcomn Reynolds", "number": "+14123444123"}, {"name": "Robbie mcKinstry", "number": "+14124453171"}]
		}
	`))
}
func HandleMemoCreate(w http.ResponseWriter, r *http.Request)     {
}


func main() {
	log.Println("Hello Memosyne")

	r := mux.NewRouter()
	r.HandleFunc("/sessions/{id}", HandleSessionRead).Methods("GET")
	r.HandleFunc("/sessions", HandleSessionCreate).Methods("POST")
	r.HandleFunc("/users/{id}/contacts", HandleContactsRead).Methods("GET")
	r.HandleFunc("/users/{id}", HandleUserRead).Methods("GET")
	r.HandleFunc("/memos", HandleMemoCreate).Methods("POST")
	r.HandleFunc("/users/{id}/memos", HandleMemoRead).Methods("GET")
	r.HandleFunc("/users", HandleUserCreate).Methods("POST")
	r.HandleFunc("/contacts", HandleContactsCreate).Methods("POST")
	r.HandleFunc("/", HandleHelloWorld)
	http.Handle("/", r)

	log.Println("Now listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HandleMemoRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["id"]
	_ = sessionId
	// get all memos out of the database for a user.
	// seralize those memos into json.
	// then write that json
	w.Write([]byte(`{
		"memos": [{"body": "Please, Jenny! I love you!"}, {"body": "I aim to misbehave!"}]
	}`))
}

func HandleContactsCreate(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}

func HandleSessionRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["id"]

	output := fmt.Sprintf("Hello, session %v", sessionId)
	_ = output

	// sample user object
	s := `{
		"user_id":      "19",
		"first_name":   "Robbie",
		"last_name":    "McKinstry",
		"phone_number": "412-445-3171",
	}`
	// what else goes in here? not contacts.
	// need /users/:id/memos
	// need /users/:id/contacts
	w.Write([]byte(s))
}

func HandleSessionCreate(w http.ResponseWriter, r *http.Request) {
	/* TODO get the json out of the request.
	{
		"user": {
			"phone_number": "412-445-3171",
			"password":     "foobar",
		}
	}
	*/
	// TODO make a new session in the db. Get back its id
	// return that id

	s := `{
		"session_id": 29  
	}`
	w.Write([]byte(s))
}

// TODO implement this so that we use the userId variable to get the struct out of the database.
func HandleUserRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	_ = userId

	s := `{
		"user": {
			"user_id": %v,
			"first_name": %v,
			"last_name": %v,
			"email_address": %v,
		}
	}`
	// s = fmt.Sprintf(s, _, _, _, _)
	w.Write([]byte(s))
}

// what do i give back after a successful User New?
// TODO determine what I should be giving back. Then give it back.
func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	s := `{
	}`
	w.Write([]byte(s))
}

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Memosyne!"))
}
