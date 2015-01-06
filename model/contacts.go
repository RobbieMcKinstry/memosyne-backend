package model

type Contact struct {
	ContactId int
	PhoneNum  string
	Status    int
	FirstName string
	LastName  string
}

type Contact_reference struct {
	contact_ref int
	contact_id  int
}

func ContactNew(user_id int, p_num string, first_name string, last_name string) *Contact {
	ret := true
	db := Db_connect()
	rows, err := db.Query("SELECT max(Contact_Reference.contact_id,Contact.cid) FROM Contact_Reference, Contact")
	if err != nil {
		ret = false
	}
	_ = ret // TODO WTF is this ret for?
	var counter int
	for rows.Next() {
		_ = rows.Scan(&counter)
	}
	counter = counter + 1
	rows, err = db.Query("INSERT INTO 'Contact_Reference' VALUES(?,?)", user_id, counter)
	//default to approved (2) status for now
	newContact := &Contact{counter, p_num, 2, first_name, last_name}
	newContact.ContactAdd()
	rows.Close()
	return newContact
}

func (this *Contact) ContactAdd() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("INSERT INTO Contact VALUES (?, ?, ?, ?, ?)", this.ContactId, this.PhoneNum, this.Status, this.FirstName, this.LastName)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

func (this *Contact) ContactDelete() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("DELETE FROM Contact WHERE Contact.cid=?", this.ContactId)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

/* Saves contact data to db */
func (this *Contact) ContactSave() {
	db := Db_connect()
	rows, err := db.Query("UPDATE Session SET phone_num=?, status=?,first_name=?,last_name=? WHERE cid=?", this.PhoneNum, this.Status, this.FirstName, this.LastName, this.ContactId)
	if err != nil {
		//Do nothing
	}
	rows.Close()
}
