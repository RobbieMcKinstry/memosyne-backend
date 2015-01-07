package model

import (
	"database/sql"
	"fmt"

	logger "github.com/Sirupsen/logrus"

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
		logger.WithFields(logger.Fields{
			"db_uri": connectionString,
			"db_error":true,
		}).Error(err)
		return nil, err
	}

	result := &ormDelegate{db}
	passed := result.CreateTablesIfNotExist()
	if !passed {
		logger.WithFields(logger.Fields{
			"db_uri": connectionString,
			"db_error":true,
		}).Error("Failed to create new tables.")
	}
	return result, nil
}

func (orm *ormDelegate) IsConnected() bool { return orm.Ping() == nil }

func (orm *ormDelegate) CreateTablesIfNotExist() bool {
	contactSQL := "CREATE TABLE IF NOT EXISTS 'Contact' ('cid' INT NOT NULL, 'phone_num' TEXT NOT NULL, 'status' INT, 'first_name' TEXT, 'last_name' TEXT, FOREIGN KEY(cid) REFERENCES Contact_Reference(contact_id))"
	userSQL := "CREATE TABLE IF NOT EXISTS 'User' ('phone_num' TEXT KEY NOT NULL,'email' TEXT NOT NULL,'first_name' TEXT,'last_name' TEXT, 'user_id' INT NOT NULL, 'password' TEXT NOT NULL, PRIMARY KEY(phone_num,email,user_id))"
	contactReferenceSQL := "CREATE TABLE IF NOT EXISTS 'Contact_Reference' ('contact_ref' INT NOT NULL, 'contact_id' INT NOT NULL, FOREIGN KEY(contact_ref) REFERENCES User(user_id),PRIMARY KEY(contact_ref,contact_id))"
	memoSQL := "CREATE TABLE IF NOT EXISTS 'Memo' ('id' INT NOT NULL, 'sender_id' INT, 'recipient_id' INT, 'body' TEXT,'time' INT NOT NULL, FOREIGN KEY(sender_id) REFERENCES User(user_id), PRIMARY KEY(id), FOREIGN KEY(recipient_id) REFERENCES Contact(cid))"
	sessionSQL := "CREATE TABLE IF NOT EXISTS 'Session' ('session_id' INT, 'expiration' INT NOT NULL, 'user_id' INT, PRIMARY KEY(session_id,user_id))"

	tables := []string{contactSQL, userSQL, contactReferenceSQL, memoSQL, sessionSQL}
	for _, tableSQL := range tables {
		result := orm.CreateTableFromString(tableSQL)
		if !result {
			logger.WithFields(logger.Fields{"orm_connected":orm.IsConnected(),"db_error":true,}).Error("Failed to intialize all of the tables.")
			return false
		}
	}
	return true
}

func (orm *ormDelegate) CreateTableFromString(creationSQL string) bool {

	stmt, err := orm.Prepare(creationSQL)
	if err != nil {
		logger.Println(err)
		return false
	}
	execInTransaction(orm, stmt)
	return true
}

func (orm *ormDelegate) newContact(contact *Contact) {
	id := orm.findIDFromTable("cid", "Contact")
	contact.ContactId = id

	stmt, err := orm.Prepare("INSERT INTO contact VALUES(?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		logger.Println(err)
		return
	}
	execInTransaction(orm, stmt, contact.ContactId, contact.PhoneNum, contact.Status, contact.FirstName, contact.LastName)
}

func (orm *ormDelegate) SaveContact(contact *Contact) *Contact {
	if contact.ContactId == 0 {
		orm.newContact(contact)
	} else {
		stmt, err := orm.Prepare("UPDATE Contact SET Contact.phone_num=?, Contact.status=?, Contact.first_name=?, Contact.last_name=? WHERE Contact.cid = ?")
		defer stmt.Close()
		if err != nil {
			logger.Println(err)
			return contact
		}
		execInTransaction(orm, stmt, contact.PhoneNum, contact.Status, contact.FirstName, contact.LastName, contact.ContactId)
	}

	return contact
}
func (orm *ormDelegate) SaveMemo(memo *Memo) *Memo {
	if memo.ID == 0 {
		orm.newMemo(memo)
	} else {
		stmt, err := orm.Prepare("UPDATE Memo SET Memo.sender_id=?, Memo.recipient_id=?, body=?, time=? WHERE Memo.id=?")
		defer stmt.Close()
		if err != nil {
			logger.Println(err)
			return memo
		}
		execInTransaction(orm, stmt, memo.SenderId, memo.RecipientId, memo.Body, memo.Time, memo.ID)
	}

	return memo

}

func (orm *ormDelegate) newMemo(memo *Memo) {
	id := orm.findIDFromTable("id", "Memo")
	memo.ID = id
	stmt, err := orm.Prepare("INSERT INTO Memo ('id', 'sender_id', 'recipient_id', 'body', 'time') VALUES (?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		logger.Println(err)
	}
	execInTransaction(orm, stmt, memo.ID, memo.SenderId, memo.RecipientId, memo.Body, memo.Time)
}

func (orm *ormDelegate) SaveUser(user *User) *User {
	if user.UserId == 0 {
		orm.newUser(user)
	} else {
		stmt, err := orm.Prepare("UPDATE User SET User.user_id=?, User.first_name=?, User.last_name=?, User.email=?, User.password=? WHERE User.phone_num=?")
		defer stmt.Close()
		if err != nil {
			logger.Println(err)
			return user
		}
		execInTransaction(orm, stmt, user.UserId, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNum)
	}
	return user
}

func (orm *ormDelegate) newUser(user *User) {
	id := orm.findIDFromTable("user_id", "User")
	user.UserId = id

	stmt, err := orm.Prepare("INSERT INTO User VALUES (?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		logger.Println(err)
		return
	}

	execInTransaction(orm, stmt, user.PhoneNum, user.Email, user.FirstName, user.LastName, user.UserId, user.Password)
}

func (orm *ormDelegate) findIDFromTable(idName, tableName string) int {
	result := 1
	rows, err := orm.Query(fmt.Sprintf("SELECT COUNT(*) FROM %v", tableName))
	if err != nil {
		logger.Println(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			logger.Println(err)
		}
	}
	if count != 0 {
		var uid int
		rows, err = orm.Query(fmt.Sprintf("SELECT MAX(%v) FROM %v", idName, tableName))
		if err != nil {
			logger.Println(err)
		}
		for rows.Next() {
			err = rows.Scan(&uid)
		}
		if err != nil {
			logger.Println(err)
		}
		result = uid + 1
	}
	return result
}

func (orm *ormDelegate) SaveSession(session *Session) *Session {
	if session.SessionId == 0 {
		session = orm.newSession(session)
	} else {
		stmt, err := orm.Prepare("UPDATE Session SET expiration=?, user_id=? WHERE session_id=?")
		defer stmt.Close()
		if err != nil {
			logger.Println(err)
			return session
		}
		execInTransaction(orm, stmt, session.Expiration, session.UserId, session.SessionId)
	}
	return session
}

func (orm *ormDelegate) newSession(session *Session) *Session {
	id := orm.findIDFromTable("session_id", "Session")
	session.SessionId = id

	stmt, err := orm.Prepare("INSERT INTO Session VALUES (?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		logger.Println(err)
		return session
	}
	execInTransaction(orm, stmt, session.SessionId, session.Expiration, session.UserId)

	return session
}

func (orm *ormDelegate) DeleteContact(contact *Contact) error { return nil }
func (orm *ormDelegate) DeleteMemo(memo *Memo) error          { return nil }
func (orm *ormDelegate) DeleteUser(user *User) error          { return nil }
func (orm *ormDelegate) DeleteSession(session *Session) error { return nil }

func execInTransaction(orm *ormDelegate, stmt *sql.Stmt, args ...interface{}) {
	var (
		err   error
		tx    *sql.Tx
		valid bool
	)
	tx, err = orm.Begin()
	valid = !rollbackOnErr(err, tx)

	_, err = tx.Stmt(stmt).Exec(args...)
	valid = valid && !rollbackOnErr(err, tx)
	if valid {
		tx.Commit()
	}
}

func rollbackOnErr(err error, tx *sql.Tx) bool {
	if err != nil {
		logger.Println(err)
		tx.Rollback()
	}
	return err != nil
}
