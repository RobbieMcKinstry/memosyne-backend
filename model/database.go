package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

type ORM interface {
	IsConnected() bool
	CreateTablesIfNotExist() bool

	SaveContact(*Contact) *Contact
	SaveMemo(*Memo) *Memo
	SaveUser(*User) *User
	SaveSession(*Session) *Session

	DeleteContact(*Contact) error
	DeleteMemo(*Memo) error
	DeleteUser(*User) error
	DeleteSession(*Session) error
}

type ormDelegate struct {
	*sql.DB
}

// TODO use the connection string...
func NewORM(connectionString string) (ORM, error) {

	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}

	result := &ormDelegate{db}
	passed := result.CreateTablesIfNotExist()
	if ! passed {
		log.Println("Failed to create new tables.")
	}
	return result, nil
}

func (orm *ormDelegate) IsConnected() bool { return orm.Ping() == nil }

func (orm *ormDelegate) CreateTablesIfNotExist() bool {
	contactSQL := "CREATE TABLE IF NOT EXISTS 'Contact' ('cid' INT NOT NULL, 'phone_num' TEXT NOT NULL, 'status' INT, 'first_name' TEXT, 'last_name' TEXT, FOREIGN KEY(cid) REFERENCES Contact_Reference(contact_id))"
	userSQL := "CREATE TABLE IF NOT EXISTS 'User' ('phone_num' TEXT KEY NOT NULL,'email' TEXT NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT NOT NULL, 'password' TEXT NOT NULL, PRIMARY KEY(phone_num,email,user_id))"
	contactReferenceSQL := "CREATE TABLE IF NOT EXISTS 'Contact_Reference' ('contact_ref' INT NOT NULL, 'contact_id' INT NOT NULL, FOREIGN KEY(contact_ref) REFERENCES User(user_id),PRIMARY KEY(contact_ref,contact_id))"
	memoSQL := "CREATE TABLE IF NOT EXISTS 'Memo' ('sender_id' INT, 'recipient_id' INT, 'body' TEXT,'time' TEXT, FOREIGN KEY(sender_id) REFERENCES User(user_id), FOREIGN KEY(recipient_id) REFERENCES Contact(cid))"
	sessionSQL := "CREATE TABLE IF NOT EXISTS 'Session' ('session_id' INT, 'expiration' TEXT, 'user_id' INT, PRIMARY KEY(session_id,user_id))"

	tables := []string{contactSQL, userSQL, contactReferenceSQL, memoSQL, sessionSQL}
	for _, tableSQL := range tables {
		result := orm.CreateTableFromString(tableSQL)
		if !result {
			return false
		}
	}
	return true
}

func (orm *ormDelegate) CreateTableFromString(creationSQL string) bool {
	rows, err := orm.Query(creationSQL)
	rows.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func (orm * ormDelegate) newContact(contact *Contact){
  id := orm.findIDFromTable("cid", "Contact")
  contact.Cid = id
  result, err := orm.Query("INSERT INTO Contact VALUES(?,?,?,?,?)", contact.Cid,contact.Phone_num,contact.Status,contact.First_name,contact.Last_name)
  if err != nil {
    log.Println(err)
  }
  result.Close()
}

func (orm *ormDelegate) SaveContact(contact *Contact) *Contact {
  if contact.Cid == 0{
    orm.newContact(contact)  
  } else{
    result,err := orm.Query("UPDATE Contact SET Contact.phone_num=?, Contact.status=?, Contact.first_name=?, Contact.last_name=? WHERE Contact.cid = ?",)
    if err != nil{
      log.Println(err)
    }
    result.Close()
  }
  
  return contact
}
func (orm *ormDelegate) SaveMemo(memo *Memo) *Memo       { return memo }
func (orm *ormDelegate) SaveUser(user *User) *User {
	if user.User_id == 0 {
		orm.newUser(user)
	} else {
		result, err := orm.Query("UPDATE User SET User.user_id=?, User.first_name=?, User.last_name=?, User.email=?, User.password=? WHERE User.phone_num=?", user.User_id, user.First_name, user.Last_name, user.Email, user.Password, user.Phone_num)
		if err != nil {
			log.Println(err)
		}
		defer result.Close()
	}
	return user
}

func (orm *ormDelegate) newUser(user *User) {
	id := orm.findIDFromTable("user_id", "User")
	user.User_id = id
	result, err := orm.Query("INSERT INTO User VALUES (?, ?, ?, ?, ?, ?)", user.Phone_num, user.Email, user.First_name, user.Last_name, user.User_id, user.Password)
	if err != nil {
		log.Println(err)
	}
	result.Close()
}

func (orm *ormDelegate) findIDFromTable(idName, tableName string) int {
	result := 1
	rows, err := orm.Query(fmt.Sprintf("SELECT COUNT(*) FROM %v", tableName))
	if err != nil {
		log.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Println(err)
		}
	}
	if count != 0 {
		var uid int
		rows, err = orm.Query(fmt.Sprintf("SELECT MAX(%v) FROM %v", idName, tableName))
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			err = rows.Scan(&uid)
		}
		if err != nil {
			log.Println(err)
		}
		result = uid + 1
	}
	return result
}

func (orm *ormDelegate) SaveSession(session *Session) *Session {
	if session.Session_id == 0 {
		session = orm.newSession(session)
	} else {
		rows, err := orm.Query("UPDATE Session SET expiration=?, user_id=? WHERE session_id=?", session.Expiration, session.User_id, session.Session_id)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
	}
	return session
}

func (orm *ormDelegate) newSession(session *Session) *Session {
	id := orm.findIDFromTable("session_id", "Session")
	session.Session_id = id

	rows, err := orm.Query("INSERT INTO Session VALUES (?, ?, ?)", session.Session_id, session.Expiration, session.User_id)
	if err != nil {
		log.Println(err)
	}
	rows.Close()
	return session
}

func (orm *ormDelegate) DeleteContact(contact *Contact) error { return nil }
func (orm *ormDelegate) DeleteMemo(memo *Memo) error          { return nil }
func (orm *ormDelegate) DeleteUser(user *User) error          { return nil }
func (orm *ormDelegate) DeleteSession(session *Session) error { return nil }
