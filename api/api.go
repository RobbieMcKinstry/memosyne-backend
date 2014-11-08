package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"time"
)

type User struct {
	phone_num   string
    email       string
	first_name  string
	last_name   string
	user_id     int
	password    string
}

type Memo struct {
	sender_id    int
	recipient_id int
	body         string
	time         string
}

type Contact struct {
	cid          int
	phone_num string
    first_name string
    last_name string
	status       int
}

type Contact_reference struct {
	contact_ref int
	contact_id  int
}

type Session struct {
	session_id int
	expiration string
	user_id    int
}

func Db_connect() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")  
	if err != nil {
		fmt.Println(err)
	}
  rows ,err2:= db.Query("PRAGMA foreign_keys = ON;")
  if err2 != nil{
    fmt.Println(err)
  }
  rows.Close()
	return db
}

//create user table
func Create_User_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'User' ('phone_num' TEXT KEY NOT NULL,'email' TEXT NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT NOT NULL, 'password' TEXT NOT NULL, PRIMARY KEY(phone_num,email,user_id))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_Contact_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'Contact' ('cid' INT NOT NULL, 'phone_num' TEXT NOT NULL, 'status' INT, FOREIGN KEY(cid) REFERENCES Contact_Reference(contact_id))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_Contact_Reference_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'Contact_Reference' ('contact_ref' INT NOT NULL, 'contact_id' INT NOT NULL, FOREIGN KEY(contact_ref) REFERENCES User(user_id),PRIMARY KEY(contact_ref,contact_id))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_Memo_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'Memo' ('sender_id' INT, 'recipient_id' INT, 'body' TEXT,'time' TEXT, FOREIGN KEY(sender_id) REFERENCES User(user_id), FOREIGN KEY(recipient_id) REFERENCES Contact(cid))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_Session_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'Session' ('session_id' INT, 'expiration' TEXT, 'user_id' INT, PRIMARY KEY(session_id,user_id))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_tables(connection *sql.DB) bool {
	if connection == nil {
		return false
	}
	if Create_User_table(connection) == false {
		return false
	}
  if Create_Contact_Reference_table(connection) == false {
		return false
	}
	if Create_Contact_table(connection) == false {
		return false
	}
  if Create_Session_table(connection) == false {
		return false
	}
	if Create_Memo_table(connection) == false {
		return false
	}
	return true
}

/*---------- GetValuesFromDBByID ----------*/

func GetSessionByID(id int, connection *sql.DB) *Session {
  rows, err := connection.Query("SELECT COUNT(*) FROM Session WHERE session_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {
    err = rows.Scan(&count)
  }
                   
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM Session WHERE session_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var s_id int
  var expr string
  var u_id int
  
  defer rows.Close()
  for rows.Next() {
    err = rows.Scan(&s_id, &expr, &u_id)
  }

  if err != nil {
    fmt.Println(err)
    return nil
  }
  
  ret := &Session{s_id, expr, u_id}
  
  return ret
}

func GetUserByPhone(phone, connection *sql.DB) *User {
  rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE phone_num=?", phone)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {
    err = rows.Scan(&count)
  }
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM User WHERE phone_num=?", phone)
  if err != nil {
    fmt.Println(err)
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
    fmt.Println(err)
    return nil
  }
  
  ret := &User{p_num, e_mail, f_name, l_name, u_id, pass}
  
  return ret
}

func GetUserByEmail(e_mail, connection *sql.DB) *User {
  rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE email=?", e_mail)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {
    err = rows.Scan(&count)
  }
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM User WHERE email=?", e_mail)
  if err != nil {
    fmt.Println(err)
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
    fmt.Println(err)
    return nil
  }
  
  ret := &User{p_num, mail, f_name, l_name, u_id, pass}
  
  return ret
}

func GetUserByID(id, connection *sql.DB) *User {
  rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE user_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {
    err = rows.Scan(&count)
  }
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM User WHERE user_id=?", id)
  if err != nil {
    fmt.Println(err)
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
    fmt.Println(err)
    return nil
  }
  
  ret := &User{p_num, mail, f_name, l_name, u_id, pass}
  
  return ret
}

func GetMemoByID(s_id int, r_id int, connection *sql.DB) *Memo {
  rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE sender_id=?, recipient_id=?", s_id, r_id)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {p
    err = rows.Scan(&count)
  }
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM User WHERE sender_id=?, recipient_id=?", s_id, r_id)
  if err != nil {
    fmt.Println(err)
  }
  
  var send_id int
  var recip_id int
  var message_body string
  var message_time string
  
  defer rows.Close()
  for rows.Next() {
    err = rows.Scan(&send_id, &recip_id, &message_body, &message_time)
  }
  
  if err != nil {
    fmt.Println(err)
    return nil
  }
  
  ret := &Memo{send_id, recip_id, message_body, message_time}
  
  return ret
}

func GetContactByID(id int, connection *sql.DB) *Contact {
  
  rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE cid=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {
    err = rows.Scan(&count)
  }
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM User WHERE cid=?", id)
  if err != nil {
    fmt.Println(err)
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
    fmt.Println(err)
    return nil
  }
  
  ret := &Contact{contact_id, p_num, f_name, l_name, stat}
  
  return ret
}

func GetContactRefByID(id int, connection *sql.DB) *Contact_reference {
  rows, err := connection.Query("SELECT COUNT(*) FROM User WHERE contact_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var count int
  for rows.Next() {
    err = rows.Scan(&count)
  }
  
  fmt.Println(count)
  
  if count == 0 {
    return nil
  }
  
  rows, err = connection.Query("SELECT * FROM User WHERE contact_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var c_ref int
  var c_id int
  
  defer rows.Close()
  for rows.Next() {
    err = rows.Scan(&c_ref, &c_id)
  }
  
  if err != nil {
    fmt.Println(err)
    return nil
  }
  
  ret := &Contact_reference{c_ref, c_id}
  
  return ret
}

func GetContactsByUserID(user_id int) []*Contacts {
    db := Db_connect()
    //contactCount, err := db.Query("SELECT COUNT(*) FROM Contact")
    contacts, err := db.Query("SELECT * FROM Contact,Contact_Reference WHERE Contact_Reference.contact_ref=? AND Contact_Reference.contact_id=Contact.cid", user_id)
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
        var theFirstName string
        var theLastName string
        var theStatus int

        err = contacts.Scan(&theCid, &thePhoneNum, &theFirstName, &theLastName, &theStatus
        )

        newContactObj := &Contact{theCid, thePhoneNum, theFirstName, theLastName, theStatus}

        contactPointerList = append(contactPointerList, newContactObj)
    }
    db.Close()
    return contactPointerList
}

/*
 a function that gets me all 
 the memos within a certain time frame
*/
func getMemosWithinRange(date1 string, date2 string) []*Memo {
  
  db = Db_connect()
  memos, err := db.Query("SELECT * FROM Memos")
  if err != nil {
    fmt.Println(err)
  }
  
  var memosWithinRange []*Memo
  memosWithinRange = make([]*Memo, len(memos))

  defer memos.Close()
  for memos.Next() {
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

      //func Parse(layout, value string) (Time, error)
      var res := Parse(Time.RFC822Z, theTime)

      //func (t Time) After(u Time) bool
      if res.Before(date1) && res.After(date2) {
        memoObj := &Memo{
          theSenderId,
          theRecipientId,
          theBody,
          theTime
        }
        memosWithinRange = append(memosWithinRange, memoObj)        
      }       
  }
  db.Close()
  return memosWithinRange
}

func main() {
  Create_tables(Db_connect())
}

