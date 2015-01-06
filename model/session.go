package model

import (
	"fmt"
	"time"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

type Session struct {
	SessionId  int
	Expiration time.Time
	UserId     int
}

func (s1 *Session) Equals(s2 *Session) bool {
	return (s1.Expiration.Equal(s2.Expiration)) && (s1.UserId == s2.UserId)
}

func (session *Session) ToString() string {
	return fmt.Sprintf("ID: %v, User ID, %v, Expiration: %v", session.SessionId, session.UserId, session.Expiration)
}

// TODO really need to abstract this out into a function that works for any kind of table
// TODO look into the Scan function for this
func (orm *ormDelegate) FindSessionByID(id int) *Session {

	var dbTime int64

	stmt, err := orm.Prepare("SELECT session_id, user_id, expiration FROM Session WHERE session_id=?")
	if err != nil {
		fmt.Println(err)
		return &Session{}
	}
	row := stmt.QueryRow(id)
	result := &Session{}
	err = row.Scan(&result.SessionId, &result.UserId, &dbTime)
	if err != nil {
		fmt.Println(err)
	}

	result.Expiration = time.Unix(dbTime, 0)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

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
		newSession = &Session{sid, time.Now().AddDate(1, 0, 0), curUser.UserId}
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
