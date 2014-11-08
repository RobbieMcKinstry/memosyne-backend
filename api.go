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

func main() {
	Create_tables(Db_connect())
}
