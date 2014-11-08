//Save changes to session
func (this *Session) save() {
  db := Db_connect()
  db.Exec(
    "INSERT INTO Session VALUES (?, ?, ?)",
    this.session_id,
    this.expiration.String(), /* String returns the time formatted using the format string */
    this.user_id
  )
}

//Make a new session
func (this *Session) newSession() *Session {
  ret := new (Session)
  return ret
}

//Delete session from db
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
  
  session_time := time.Parse(time.RFC822Z, this.expiration)
  
  if session_time.After(time.Now()) == true {
    ret = true
  }
  
  return ret
}