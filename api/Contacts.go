/*--------- Contacts Methods ----------*/
func ContactNew(contact_id int, p_num string, stat int) *Contact {
  newContact := &Contact{contact_id, p_num, stat}
  newContact.ContactAdd()
  return newContact
}

func (this *Contact) ContactAdd() bool {
  ret := true
  db := Db_connect()
  
  rows, err := db.Query("INSERT INTO Contact VALUES (?, ?, ?)", this.cid, this.phone_num, this.status)
  if err != nil {
    ret = false
  }
  rows.Close()
  
  return ret
}

func (this *Contact) ContactDelete() bool {
  ret := true
  db := Db_connect()

  rows, err := db.Query("DELETE FROM Contact WHERE Contact.cid=?", this.cid)
  if err != nil {
      ret = false
  }
  rows.Close()
  
  return ret
}

/* Saves contact data to db */
func (this *Contact) ContactSave() {
  db := Db_connect()
  rows, err := db.Query("UPDATE Session SET phone_num=?, status=? WHERE cid=?", this.phone_num, this.status, this.cid)
  if err != nil {
    //Do nothing
  }
  rows.Close()
}
/*--------- END CONTACTS -----------*/