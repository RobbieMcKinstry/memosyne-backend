package model

import (
	"time"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

func MemoNew(id, sid, rid int, b string, t time.Time) *Memo {
	newMemo := &Memo{id, sid, rid, b, t}

	db := Db_connect()
	rows, _ := db.Query("SELECT Contact.status FROM Contact, Contact_Reference WHERE Contact_Reference.contact_ref = ? AND Contact.cid = ? AND Contact.cid = Contact_Reference.contact_id;", sid, rid)

	var contactstatus int

	for rows.Next() {
		_ = rows.Scan(&contactstatus)
	}
	if contactstatus == 2 {
		newMemo = &Memo{id, sid, rid, b, t}
		newMemo.MemoAdd()
	} else {
		newMemo = nil
	}
	return newMemo
}

/* Adds Memo to Database */
func (this *Memo) MemoAdd() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("INSERT INTO Memo VALUES (?, ?, ?, ?)", this.SenderId, this.RecipientId, this.Body, this.Time)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}

/* Saves contact data to db */
func (this *Memo) MemoSave() {
	db := Db_connect()
	rows, err := db.Query("UPDATE Memo SET body=?, time=? WHERE sender_id=? AND recipient_id", this.Body, this.Time, this.SenderId, this.RecipientId)
	if err != nil {
		//Do nothing
	}
	rows.Close()
}

func (this *Memo) MemoDelete() bool {
	ret := true
	db := Db_connect()

	rows, err := db.Query("DELETE FROM Memo WHERE Memo.sender_id=? AND Memo.recipient_id=?", this.SenderId, this.RecipientId)
	if err != nil {
		ret = false
	}
	rows.Close()

	return ret
}
