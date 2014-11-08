/*---------- Memos Methods ----------*/
func MemoNew(sid int, rid int, b string, t string) *Memo {
  newMemo := &Memo{sid, rid, b, t}
  
  db := Db_connect()
  rows, err := db.Query("SELECT Contact.status FROM Contact, Contact_Reference WHERE Contact_Reference.contact_ref = ? AND Contact.cid = ? AND Contact.cid = Contact_Reference.contact_id;", sid,rid)
  
  var contactstatus int;
  
  for rows.Next(){
    err := rows.Scan(&contactstatus)
  }
  if(contactstatus == 2){
    newMemo := {sid,rid,b,t}
    newMemo.MemoAdd()
  } else {
    newMemo := nil
  }
  return newMemo
}

/* Adds Memo to Database */
func (this *Memo) MemoAdd() bool {
  ret := true
  db := Db_connect()
  
  rows, err := db.Query("INSERT INTO Memo VALUES (?, ?, ?, ?)", this.sender_id, this.recipient_id, this.body, this.time)
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
/*---------- MEMOS END ----------*/