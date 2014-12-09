package model

import (
	"fmt"
	"time"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

type Memo struct {
	ID          int
	SenderId    int
	RecipientId int
	Body        string
	Time        time.Time
}

func (memo *Memo) Equals(other *Memo) bool {
	return memo.ID == other.ID &&
		memo.SenderId == other.SenderId &&
		memo.RecipientId == other.RecipientId &&
		memo.Body == other.Body &&
		memo.Time.Equal(other.Time)
}

func (orm *ormDelegate) FindMemoByID(id int) *Memo {
	var dbTime int64

	stmt, err := orm.Prepare("SELECT id, sender_id, recipient_id, body, time FROM Memo WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return &Memo{}
	}
	row := stmt.QueryRow(id)
	result := &Memo{}
	err = row.Scan(&result.ID, &result.SenderId, &result.RecipientId, &result.Body, &dbTime)
	if err != nil {
		fmt.Println(err)
		return &Memo{}
	}
	result.Time = time.Unix(dbTime, 0)
	return result
}

func (memo *Memo) ToString() string {
	return fmt.Sprintf("Memo [ID: %v, Sender: %v, Recipient: %v, Time: %v, Body: %v ]", memo.ID, memo.SenderId, memo.RecipientId, memo.Time.Format(time.RFC822), memo.Body)
}

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
