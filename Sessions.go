/*---------- Session Methods ----------*/

//Save changes to session
func (this *Session) Save() {
  db := Db_connect()
  rows, err := db.Query("UPDATE Session SET expiration=?, user_id=? WHERE session_id=?", this.expiration, this.user_id, this.session_id)
  if err != nil {
    //Do nothing
  }
  rows.Close()
}

//Make a new session
func (this *Session) SessionNew() bool {
  ret := true
  db := Db_connect()
  
  rows, err := db.Query("INSERT INTO Session VALUES (?, ?, ?)", this.session_id, this.expiration, this.user_id)
  if err != nil {
    ret = false
  }
  rows.Close()
  
  return ret
}

//Delete session from db
func (this *Session) Delete() bool {
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