package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

var log = logrus.New()

func Db_connect() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Println(err)
	}
	rows, err2 := db.Query("PRAGMA foreign_keys = ON;")
	if err2 != nil {
		log.Println(err)
	}
	rows.Close()
	return db
}

func GetSessionByID(id int, connection *sql.DB) *Session {
	rows, err := connection.Query("SELECT COUNT(*) FROM Session WHERE session_id=?", id)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM Session WHERE session_id=?", id)
	if err != nil {
		log.Println(err)
	}

	var s_id int
	var expr time.Time
	var u_id int

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&s_id, &expr, &u_id)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &Session{s_id, expr, u_id}

	return ret
}

func GetUserByPhone(phone, connection *sql.DB) *User {
	rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE phone_num=?", phone)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM User WHERE phone_num=?", phone)
	if err != nil {
		log.Println(err)
	}

	var p_num string
	var e_mail string
	var f_name string
	var l_name string
	var u_id int
	var pass string

	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&p_num, &e_mail, &f_name, &l_name, &u_id, &pass)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &User{p_num, e_mail, f_name, l_name, u_id, pass}

	return ret
}

func GetUserByEmail(connection *sql.DB, e_mail string) *User {
	rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE email=?", e_mail)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM User WHERE email=?", e_mail)
	if err != nil {
		log.Println(err)
	}

	var p_num string
	var mail string
	var f_name string
	var l_name string
	var u_id int
	var pass string

	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&p_num, &mail, &f_name, &l_name, &u_id, &pass)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &User{p_num, mail, f_name, l_name, u_id, pass}

	return ret
}

func GetUserByID(id, connection *sql.DB) *User {
	rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE user_id=?", id)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM User WHERE user_id=?", id)
	if err != nil {
		log.Println(err)
	}

	var p_num string
	var mail string
	var f_name string
	var l_name string
	var u_id int
	var pass string

	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&p_num, &mail, &f_name, &l_name, &u_id, &pass)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &User{p_num, mail, f_name, l_name, u_id, pass}

	return ret
}

func GetMemoByID(s_id int, r_id int, connection *sql.DB) *Memo {
	rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE sender_id=?, recipient_id=?", s_id, r_id)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM User WHERE sender_id=?, recipient_id=?", s_id, r_id)
	if err != nil {
		log.Println(err)
	}

	var id int
	var send_id int
	var recip_id int
	var message_body string
	var message_time time.Time

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &send_id, &recip_id, &message_body, &message_time)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &Memo{id, send_id, recip_id, message_body, message_time}

	return ret
}

func GetContactByID(id int, connection *sql.DB) *Contact {

	rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE cid=?", id)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM User WHERE cid=?", id)
	if err != nil {
		log.Println(err)
	}

	var contact_id int
	var p_num string
	var f_name string
	var l_name string
	var stat int

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&contact_id, &p_num, &f_name, &l_name, &stat)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	//ret := &Contact{contact_id, p_num, f_name, l_name, stat}

	return new(Contact) // TODO implement a patch!
}

func GetContactRefByID(id int, connection *sql.DB) *Contact_reference {
	rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE contact_id=?", id)
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	log.Println(count)

	if count == 0 {
		return nil
	}

	rows, err = connection.Query("SELECT * FROM User WHERE contact_id=?", id)
	if err != nil {
		log.Println(err)
	}

	var c_ref int
	var c_id int

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&c_ref, &c_id)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &Contact_reference{c_ref, c_id}

	return ret
}

func GetContactsByUserID(user_id int) []*Contact {
	db := Db_connect()
	//contactCount, err := db.Query("SELECT COUNT(*) FROM Contact")
	contacts, err := db.Query("SELECT * FROM Contact,Contact_Reference WHERE Contact_Reference.contact_ref=? AND Contact_Reference.contact_id=Contact.cid", user_id)
	if err != nil {
		log.Println(err)
	}

	var contactPointerList []*Contact
	contactPointerList = make([]*Contact, 0)

	defer contacts.Close()
	for contacts.Next() {
		newContactObj := new(Contact)

		var theCid int
		var thePhoneNum string
		var theFirstName string
		var theLastName string
		var theStatus int

		err = contacts.Scan(&theCid, &thePhoneNum, &theFirstName, &theLastName, &theStatus)

		//newContactObj = &Contact{theCid, thePhoneNum, theFirstName, theLastName, theStatus} // TODO implement a patch

		contactPointerList = append(contactPointerList, newContactObj)
	}
	db.Close()
	return contactPointerList
}

/*
 a function that gets me all
 the memos within a certain time frame
*/
func GetMemosWithinRange(date1 string, date2 string) []*Memo {

	db := Db_connect()
	memos, err := db.Query("SELECT * FROM Memos")
	if err != nil {
		log.Println(err)
	}

	var memosWithinRange []*Memo
	memosWithinRange = make([]*Memo, 0)

	defer memos.Close()
	for memos.Next() {
		var theSenderId int
		var theRecipientId int
		var theBody string
		var theTime time.Time

		_ = memos.Scan(
			&theSenderId,
			&theRecipientId,
			&theBody,
			&theTime,
		)

		//func Parse(layout, value string) (Time, error)
		//res := Parse(Time.RFC822Z, theTime)

		//func (t Time) After(u Time) bool
		//if res.Before(date1) && res.After(date2) {
		//	memoObj := &Memo{
		//		theSenderId,
		//		theRecipientId,
		//		theBody,
		//		theTime,
		//	}
		//	memosWithinRange = append(memosWithinRange, memoObj)
		//}
	}
	db.Close()
	return memosWithinRange
}

/* returns a dynamically allocated slice containing pointers to Memo objects related to the userID passed */
func GetMemosByUserID(uid int) []*Memo {
	db := Db_connect()

	memos, err := db.Query("SELECT * FROM Memos")
	if err != nil {
		log.Println(err)
	}

	var memoPointerList []*Memo
	memoPointerList = make([]*Memo, 0)

	defer memos.Close()
	for memos.Next() {

		var (
			id             int
			theSenderId    int
			theRecipientId int
			theBody        string
			theTime        time.Time
		)

		err = memos.Scan(&id, &theSenderId, &theRecipientId, &theBody, &theTime)

		if theSenderId == uid {
			newMemoObj := &Memo{id, theSenderId, theRecipientId, theBody, theTime}
			memoPointerList = append(memoPointerList, newMemoObj)
		}
	}
	db.Close()
	return memoPointerList
}

/* returns a dynamically allocated slice containing pointers to Memo objects */
func GetMemos() []*Memo {
	db := Db_connect()

	memos, err := db.Query("SELECT * FROM Memos")
	if err != nil {
		log.Println(err)
	}

	var memoPointerList []*Memo
	memoPointerList = make([]*Memo, 0)

	defer memos.Close()
	for memos.Next() {

		var (
			id             int
			theSenderId    int
			theRecipientId int
			theBody        string
			theTime        time.Time
		)

		err = memos.Scan(&id, &theSenderId, &theRecipientId, &theBody, &theTime)

		newMemoObj := &Memo{id, theSenderId, theRecipientId, theBody, theTime}
		memoPointerList = append(memoPointerList, newMemoObj)
	}

	db.Close()
	return memoPointerList
}
