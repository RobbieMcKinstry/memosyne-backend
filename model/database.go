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

	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return nil, err
	}
	CreateTables(db)

	result := &ormDelegate{db}
	return result, nil
}

func (orm *ormDelegate) IsConnected() bool { return orm.Ping() == nil }

func (orm *ormDelegate) CreateTablesIfNotExist() bool {
	contactSQL := "CREATE TABLE IF NOT EXISTS 'Contact' ('cid' INT NOT NULL, 'phone_num' TEXT NOT NULL, 'status' INT, FOREIGN KEY(cid) REFERENCES Contact_Reference(contact_id))"
	result := orm.CreateTableFromString(contactSQL)
	if !result {
		return false
	}

	userSQL := "CREATE TABLE IF NOT EXISTS 'User' ('phone_num' TEXT KEY NOT NULL,'email' TEXT NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT NOT NULL, 'password' TEXT NOT NULL, PRIMARY KEY(phone_num,email,user_id))"
	result = orm.CreateTableFromString(userSQL)
	if !result {
		return false
	}

	contactReferenceSQL := "CREATE TABLE IF NOT EXISTS 'Contact_Reference' ('contact_ref' INT NOT NULL, 'contact_id' INT NOT NULL, FOREIGN KEY(contact_ref) REFERENCES User(user_id),PRIMARY KEY(contact_ref,contact_id))"
	result = orm.CreateTableFromString(contactReferenceSQL)
	if !result {
		return false
	}

	memoSQL := "CREATE TABLE IF NOT EXISTS 'Memo' ('sender_id' INT, 'recipient_id' INT, 'body' TEXT,'time' TEXT, FOREIGN KEY(sender_id) REFERENCES User(user_id), FOREIGN KEY(recipient_id) REFERENCES Contact(cid))"
	result = orm.CreateTableFromString(memoSQL)
	if !result {
		return false
	}

	sessionSQL := "CREATE TABLE IF NOT EXISTS 'Session' ('session_id' INT, 'expiration' TEXT, 'user_id' INT, PRIMARY KEY(session_id,user_id))"
	result = orm.CreateTableFromString(sessionSQL)
	if !result {
		return false
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

func (orm *ormDelegate) SaveContact(c *Contact) *Contact { return c }
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
	id := orm.findIDFromTable("User")
	user.User_id = id
	result, err := orm.Query("INSERT INTO User VALUES (?, ?, ?, ?, ?, ?)", user.Phone_num, user.Email, user.First_name, user.Last_name, user.User_id, user.Password)
	if err != nil {
		log.Println(err)
	}
	result.Close()
}

func (orm *ormDelegate) findIDFromTable(tableName string) int {
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
		rows, err = orm.Query("SELECT MAX(user_id) FROM User")
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

func (orm *ormDelegate) SaveSession(session *Session) *Session { return session }

func (orm *ormDelegate) DeleteContact(contact *Contact) error { return nil }
func (orm *ormDelegate) DeleteMemo(memo *Memo) error          { return nil }
func (orm *ormDelegate) DeleteUser(user *User) error          { return nil }
func (orm *ormDelegate) DeleteSession(session *Session) error { return nil }
