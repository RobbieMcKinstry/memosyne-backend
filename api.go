package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"time"
)

type User struct {
	phone_num   string
	first_name  string
	last_name   string
	user_id     int
	contact_ref int
	password    string
}

type Memo struct {
	sender_id    int
	recipient_id int
	body         string
	time         time.Time
}

type Contact struct {
	cid          int
	phone_number string
	status       int
}

type Contact_reference struct {
	contact_ref int
	contact_id  int
}

type Session struct {
	session_id int
	expiration time.Time
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
	createString := "CREATE TABLE IF NOT EXISTS 'User' ('phone_num' TEXT KEY NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT PRIMARY KEY NOT NULL,'contact_ref' INT NOT NULL, 'password' TEXT NOT NULL)"
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
	createString := "CREATE TABLE IF NOT EXISTS 'Contact_Reference' ('contact_ref' INT NOT NULL, 'contact_id' INT PRIMARY KEY NOT NULL, FOREIGN KEY(contact_ref) REFERENCES User(contact_ref))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_Memo_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'Memo' ('sender_id' INT, 'recipient_id' INT, 'body' TEXT,'time' datetime, FOREIGN KEY(sender_id) REFERENCES User(user_id), FOREIGN KEY(recipient_id) REFERENCES COntact(cid))"
	rows, err := connection.Query(createString)
	rows.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Create_Session_table(connection *sql.DB) bool {
	createString := "CREATE TABLE IF NOT EXISTS 'Session' ('session_id' INT, 'expiration' datetime, 'user_id' INT, FOREIGN KEY(user_id) REFERENCES User(user_id))"
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

func GetSessionByID(id int, connection *sql.DB) *Session {
  rows, err := connection.Query("SELECT * FROM Session WHERE session_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var s_id int
  var expr time.Time
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

func GetUserByID(id int, connection *sql.DB) *User {
  rows, err := connection.Query("SELECT * FROM User WHERE user_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var p_num string
  var f_name string
  var l_name string
  var u_id int
  var c_ref int
  var pass string
  
  defer rows.Close()
  for rows.Next() {

    err = rows.Scan(&p_num, &f_name, &l_name, &u_id, &c_ref, &pass)
  }
    
  if err != nil {
    fmt.Println(err)
    return nil
  }
  
  ret := &User{p_num, f_name, l_name, u_id, c_ref, pass}
  
  return ret
}

func GetMemoByID(s_id int, r_id int, connection *sql.DB) *Memo {
  rows, err := connection.Query("SELECT * FROM User WHERE sender_id=?, recipient_id=?", s_id, r_id)
  if err != nil {
    fmt.Println(err)
  }
  
  var send_id int
  var recip_id int
  var message_body string
  var message_time time.Time
  
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
  rows, err := connection.Query("SELECT * FROM User WHERE c_id=?", id)
  if err != nil {
    fmt.Println(err)
  }
  
  var contact_id int
  var p_num string
  var stat int
  
  defer rows.Close()
  for rows.Next() {
    err = rows.Scan(&contact_id, &p_num, &stat)
  }
  
  if err != nil {
    fmt.Println(err)
    return nil
  }
  
  ret := &Contact{contact_id, p_num, stat}
  
  return ret
}

func GetContactRefByID(id int, connection *sql.DB) *Contact_reference {
  rows, err := connection.Query("SELECT * FROM User WHERE contact_id=?", id)
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

func main() {
	
}
