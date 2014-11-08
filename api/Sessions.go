/*---------- Session Methods ----------*/

//Make new session
func SessionNew(sid int, expr string, uid int) *Session {
  newSession := &Session{sid, expr, uid}
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