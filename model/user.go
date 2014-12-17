package model

type User struct {
	PhoneNum  string
	Email     string
	FirstName string
	LastName  string
	UserId    int
	Password  string
}

/* Makes new User */
func UserNew(p_num int, email string, f_name string, l_name string, pass string) *User {
	db := Db_connect()
	var newUser *User
	var uid int
	rows, err := db.Query("SELECT MAX(user_id) FROM User")
	if err != nil {
		newUser = nil
	}
	for rows.Next() {
		err = rows.Scan(&uid)
	}
	if err != nil {
		newUser = nil
	}
	uid = uid + 1
	//newUser = &User{p_num, email, f_name, l_name, uid, pass}
	newUser.UserAdd()
	return newUser
}

/* INSERTS values in db */
func (this *User) UserAdd() bool {
	ret := true
	db := Db_connect()

	result, err := db.Query("INSERT INTO User VALUES (?, ?, ?, ?, ?, ?)", this.PhoneNum, this.Email, this.FirstName, this.LastName, this.UserId, this.Password)
	result.Close()
	if err != nil {
		log.Println(err)
		ret = false
	}

	db.Close()
	return ret
}

/* UPDATES values in db */
func (this *User) UserSave() {
	db := Db_connect()

	result, err := db.Query("UPDATE User SET User.user_id=?, User.first_name=?, User.last_name=?, User.email=?, User.password=? WHERE User.phone_num=?", this.UserId, this.FirstName, this.LastName, this.Email, this.Password, this.PhoneNum)
	result.Close()
	if err != nil {
		//Do nothing
	}
	db.Close()
}

func (this *User) UserDelete() bool {
	ret := true
	db := Db_connect()
	result, err := db.Query("DELETE FROM User WHERE User.phone_num=?", this.PhoneNum)
	if err != nil {
		ret = false
	}
	result.Close()
	db.Close()
	return ret
}

func (this *User) equals(externalUser *User) bool {
	ret := false
	if externalUser.PhoneNum == this.PhoneNum &&
		externalUser.FirstName == this.FirstName &&
		externalUser.LastName == this.LastName &&
		externalUser.UserId == this.UserId &&
		externalUser.Password == this.Password &&
		externalUser.Email == this.Email {
		ret = true
	}
	return ret
}

/*
 * Returns a slice of pointers to all Contact objects stored in the db
 *
func (this *User) GetContacts() []*Contact {
    db := Db_connect()
    //contactCount, err := db.Query("SELECT COUNT(*) FROM Contact")
    contacts, err := db.Query("SELECT * FROM Contact")
    if err != nil {
        log.Println(err)
    }

    var contactPointerList []*Contact
    contactPointerList = make([]*Contact, len(contacts))

    defer contacts.Close()
    for contacts.Next() {
        newContactObj := new (Contact)

        var theCid int
        var thePhoneNum string
        var theStatus int

        err = contacts.Scan(
            &theCid,
            &thePhoneNum,
            &theStatus
        )

        newContactObj.cid = theCid
        newContactObj.phone_num = thePhoneNum
        newContactObj.status = theStatus

        contactPointerList = append(contactPointerList, newContactObj)
    }
    db.Close()
    return contactPointerList
}

/* returns a dynamically allocated slice containing pointers to Memo objects
func (this *User) GetMemos() []*Memos {
    db := Db_connect()

    memos, err := db.Query("SELECT * FROM Memos")
    if err != nil {
        log.Println(err)
    }

    var memoPointerList []*Memos
    memoPointerList = make([]*Memo, len(memos))

    defer memos.Close()
    for memos.Next() {
        newMemoObj := new (Memo)

        var theSenderId int
        var theRecipientId int
        var theBody string
        var theTime time.Time

        err = memos.Scan(
            &theSenderId,
            &theRecipientId,
            &theBody,
            &theTime
        )

        newMemoObj.senderId = theSenderId
        newMemoObj.recipientId = theRecipientId
        newMemoObj.body = theBody
        newMemoObj.time = theTime

        memoPointerList = append(memoPointerList, newMemoObj)
    }
    db.Close()
    return memoPointerList
}*/
