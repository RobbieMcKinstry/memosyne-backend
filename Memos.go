/*---------- Memos Methods ----------*/
func (this *Memo) MemoNew() bool {
  ret := true
  db := Db_connect()
  
  rows, err := db.Query("INSERT INTO Memo VALUES (?, ?, ?, ?)", this.sender_id, this.recipient_id, this.body, this.time)
  if err != nil {
    ret = false
  }
  rows.Close()
  
  return ret
}

func (this *Memo) Delete() bool {
  ret := true
  db := Db_connect()

  rows, err := db.Query("DELETE FROM Memo WHERE Memo.sender_id=? AND Memo.recipient_id=?", this.sender_id, this.recipient_id)
  if err != nil {
      ret = false
  }
  rows.Close()
  
  return ret
}

/* Saves contact data to db */
func (this *Memo) Save() {
  db := Db_connect()
  rows, err := db.Query("UPDATE Memo SET body=?, time=? WHERE sender_id=? AND recipient_id", this.body, this.time, this.sender_id, this.recipient_id)
  if err != nil {
    //Do nothing
  }
  rows.Close()
}
/*---------- MEMOS END ----------*/