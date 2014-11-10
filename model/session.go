package session

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

/*---------- Session Methods ----------*/

//Make new session
func SessionNew(email string, pass string) *Session {
	db := Db_connect()
	var newSession *Session
	var sid int
	curUser := GetUserByEmail(email, db)
	if pass != curUser.password {
		//return error
	} else {
		//generate session id
		rows, err := db.Query("SELECT MAX(session_id) FROM Session")
		for rows.Next() {
			err := rows.Scan(&sid)
		}
		if err != nil {
			newSession = nil
		}
		sid = sid + 1
		newSession = &Session{sid, "02 Jan 15 15:04 -0700", curUser.user_id}
	}
	newSession.SessionAdd()
	return newSession
}

//Add session to database
func (this *Session) SessionAdd() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("INSERT INTO Session VALUES (?, ?, ?)", this.session_id, this.expiration, this.user_id)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

//Save changes to session
func (this *Session) SessionSave() {
	db := Db_connect()
	rows, err := db.Query("UPDATE Session SET expiration=?, user_id=? WHERE session_id=?", this.expiration, this.user_id, this.session_id)
	if err != nil {
		//Do nothing
	}
	rows.Close()
}

//Delete session from db
func (this *Session) SessionDelete() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("DELETE FROM Session WHERE Session.session_id=?", this.session_id)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

func (this *Session) IsValid() bool {
	ret := false

	session_time, _ := time.Parse(time.RFC822Z, this.expiration)

	if session_time.After(time.Now()) == true {
		ret = true
	}

	return ret
}

/*---------- END SESSIONS ----------*/
