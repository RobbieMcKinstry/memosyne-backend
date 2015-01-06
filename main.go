package main

import (
	logger "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"./model"

	"fmt"
	"net/http"
	"strconv"
)

func main() {
	logger.Println("Hello Memosyne")

	r := mux.NewRouter()
	r.HandleFunc("/sessions/{id}", HandleSessionRead).Methods("GET")
	r.HandleFunc("/sessions", HandleSessionCreate).Methods("POST")

	r.HandleFunc("/users/{id}/contacts", HandleContactsRead).Methods("GET")
	r.HandleFunc("/users/{id}/memos", HandleMemoRead).Methods("GET")
	r.HandleFunc("/users/{id}", HandleUserRead).Methods("GET")
	r.HandleFunc("/users", HandleUserCreate).Methods("POST")

	r.HandleFunc("/memos", HandleMemoCreate).Methods("POST")

	r.HandleFunc("/contacts", HandleContactsCreate).Methods("POST")
	r.HandleFunc("/", HandleHelloWorld)
	http.Handle("/", r)

	logger.Println("Now listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HandleSessionRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID, _ := strconv.Atoi(vars["id"])

	orm, err := model.NewORM("development.db")
	if err != nil {
		logger.Println(err)
	}
	var session *model.Session = orm.FindSessionByID(sessionID)
	_ = session

	s := fmt.Sprintf(`{ "session_id": %v, }`, session)

	w.Write([]byte(s))
}

func HandleSessionCreate(w http.ResponseWriter, r *http.Request) {
	s := `{
		"session_id": 29,
		"user_id":    120,
	}`
	w.Write([]byte(s))
}

func HandleContactsRead(w http.ResponseWriter, r *http.Request) {
	s := `{
		“contacts”: [ 
			{ 
				“contact_id”: 1,
				“phone_number”: “412-445-3171”,
				“first_name”:	“Robbie”,
				“last_name”:	“McKinstry”, 
			}, {
				“contact_id”: 2,
				“phone_number”: “412-661-2963”,
				“first_name”:	“Malcolm”,
				“last_name”:	“Reynolds”, 
			}
		],
	}`
	w.Write([]byte(s))
}

func HandleMemoRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["id"]
	_ = sessionId
	// get all memos out of the database for a user.
	// seralize those memos into json.
	// then write that json
	s := `{

		“sender_id”:	0, 		
		“contact_id”:	0, 		
		“body”:		“”, 		
		“time”:		“”,

		"memos": [
			{
				"sender_id": 7,
				"contact_id": 99,
				"body": "Please, Jenny! I love you!",
				"time": "2006-01-02T15:04:05Z07:00",
			}, {
				"sender_id": 8,
				"contact_id": 100,
				"body": "I aim to misbehave!"
				"time": "2006-01-02T15:04:05Z07:00",
			},
		]
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
			"user_id": 12,
			"first_name": "Nigel",
			"last_name": "Thornberry",
			"email_address": "wounderous_animals@thornberry.tv",
		},
	}`
	w.Write([]byte(s))
}

// what do i give back after a successful User New?
// TODO determine what I should be giving back. Then give it back.
func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	s := `{
		"session_id": 13,
		"user_id":    201,
	}`
	w.Write([]byte(s))
}

func HandleMemoCreate(w http.ResponseWriter, r *http.Request) {
	s := ` {
		"memos": {
			"sender_id": 7,
			"contact_id": 99,
			"body": "Please, Jenny! I love you!",
			"time": "2006-01-02T15:04:05Z07:00",
		},
	}`
	w.Write([]byte(s))
}

func HandleContactsCreate(w http.ResponseWriter, r *http.Request) {
	s := `{
		“contact”: { 
			“contact_id”: 3,
			“phone_number”: “724-903-3112”,
			“first_name”:	“Robert”,
			“last_name”:	“Floyd”,
		},
	}`

	w.Write([]byte(s))
}

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Memosyne!"))
}
