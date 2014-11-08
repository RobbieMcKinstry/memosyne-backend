/*
 * Methods for Session class
 *      Author: Tyler Raborn
 */

 package main

 import (
    "database/sql"
    "fmt"
    _ "github.com/mxk/go-sqlite/sqlite3"
    "time"
)

func (this *Session) save() {
    db := Db_connect()
    db.Exec(
        "INSERT INTO Session VALUES (?, ?, ?)", 
        this.session_id,
        this.expiration.String(), /* String returns the time formatted using the format string */
        this.user_id
    )
    db.Close()
}

func (this *Session) newSession() *Session {
    ret := new (Session)
    return ret
}

func (this *Session) delete() bool {
    ret := true
    db := Db_connect()
    result, err := db.Query("DELETE FROM Session WHERE Session.session_id=?", this.session_id)
    if err != nil {
        ret = false
    }
    return ret
}

func (this *Session) isValid() bool {
    ret := false
    db := Db_connect()

    result, err := db.Query("SELECT expiration FROM Session WHERE Session.session_id=?", this.session_id)
    if err != nil {
        fmt.Println(err)
    }

    if Time.Parse(result).After(Time.Now()) {
        ret = true
    }

    db.Close()
    return ret
}
