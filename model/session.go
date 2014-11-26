package model

import (
	"time"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

//Make new session
func SessionNew(email string, pass string) *Session {
	db := Db_connect()
	var newSession *Session
	var sid int
	curUser := GetUserByEmail(db, email)
	if pass != curUser.Password {
		//return error
	} else {
		//generate session id
		rows, err := db.Query("SELECT MAX(session_id) FROM Session")
		for rows.Next() {
			_ = rows.Scan(&sid)
		}
		if err != nil {
			newSession = nil
		}
		sid = sid + 1
		newSession = &Session{sid, time.Now(), curUser.UserId}
	}
	newSession.SessionAdd()
	return newSession
}

//Add session to database
func (this *Session) SessionAdd() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("INSERT INTO Session VALUES (?, ?, ?)", this.SessionId, this.Expiration, this.UserId)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

//Save changes to session
func (this *Session) SessionSave() {
	db := Db_connect()
	rows, err := db.Query("UPDATE Session SET expiration=?, user_id=? WHERE session_id=?", this.Expiration, this.UserId, this.SessionId)
	if err != nil {
		//Do nothing
	}
	rows.Close()
}

//Delete session from db
func (this *Session) SessionDelete() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("DELETE FROM Session WHERE Session.session_id=?", this.SessionId)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

func (this *Session) IsValid() bool {
	ret := false

	session_time := this.Expiration

	if session_time.After(time.Now()) == true {
		ret = true
	}

	return ret
}
