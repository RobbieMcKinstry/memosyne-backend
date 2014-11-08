/* 
 *  Methods for the User class
 *      Written by Tyler Raborn
 */

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mxk/go-sqlite/sqlite3"
    "time"
)

/* UPDATES values in db */
func (this *User) UserSave() {
    db := Db_connect()

    result, err := db.Query(
        "UPDATE User 
            SET User.phone_num=? AND
            SET User.first_name=? AND
            SET User.last_name=? AND
            SET User.contact_ret=? AND
            SET User.password=?
        WHERE User.user_id=?",
        this.phone_num,
        this.first_name,
        this.last_name,
        this.contact_ret,
        this.password,
        this.user_id
    )
    result.Close()
    db.Close()
}

/* INSERTS values in db */
func (this *User) UserNew() bool {
    ret := true
    db := Db_connect()

    result, err := db.Query(
        "INSERT INTO User VALUES (?, ?, ?, ?, ?, ?)",
        this.phone_num,
        this.first_name,
        this.last_name,
        this.user_id,
        this.contact_ret,
        this.password
    )
    result.Close()
    if err != nil {
        fmt.Println(err)
        ret = false
    }

    db.Close()
    return ret
}

func (this *User) UserDelete() bool {
    ret := true
    db := Db_connect()
    result, err := db.Query("DELETE FROM User WHERE User.user_id=?", this.user_id)
    if err != nil {
        ret = false
    }
    result.Close()
    db.Close()
    return ret
}   

func (this *User) equals(externalUser *User) bool {
    ret := false
    if externalUser.phone_num == this.phone_num &&
       externalUser.first_name == this.first_name &&
       externalUser.last_name == this.last_name &&
       externalUser.user_id == this.user_id &&
       externalUser.contact_ret == this.contact_ret &&
       externalUser.password == this.password {
         ret = true
       }
    return ret
}

/*
 * Returns a slice of pointers to all Contact objects stored in the db
 */
func (this *User) GetContacts() []*Contact {
    db := Db_connect()
    //contactCount, err := db.Query("SELECT COUNT(*) FROM Contact")
    contacts, err := db.Query("SELECT * FROM Contact")
    if err != nil {
        fmt.Println(err)
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

/* returns a dynamically allocated slice containing pointers to Memo objects */
func (this *User) GetMemos() []*Memos { 
    db := Db_connect()

    memos, err := db.Query("SELECT * FROM Memos")
    if err != nil {
        fmt.Println(err)
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
}

