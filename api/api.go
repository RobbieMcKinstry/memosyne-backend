package api

import (
  "fmt"
  "time"
  _ "github.com/mxk/go-sqlite/sqlite3"
  "database/sql"
)

type Memo struct {
  sender_id int
  recipient_id int
  body string
  time time.Time
}

type Contact struct {
  cid int
  phone_number string
  status int
}

type Contact_reference struct {
  contact_ref int
  contact_id int
}

type Session struct {
  session_id int
  expiration time.Time
  user_id int
}


/*func Db_connect() *sql.DB {
  db, err := sql.Open("sqlite.db")
  if err != nil {
    fmt.Println(err)
  }
  return db
}

func Create_tables(connection *sql.DB) {
  //dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"SQLite", "UTF8"}}
  
  //create the user table
  createString := "CREATE TABLE IF NOT EXISTS ('phone_num' INT PRIMARY KEY NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT PRIMARY KEY NOT NULL,'contact_ref' INT PRIMARY KEY NOT NULL, 'password' TEXT PRIMARY KEY NOT NULL);"
  rowsconnection.Query(createString)
}

func main() {
  Create_tables
}*/