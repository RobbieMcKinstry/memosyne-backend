package model

import (
	"database/sql"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

type ORM interface {
	IsConnected() bool

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

func NewORM(connectionString string) (ORM, error) {

	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return nil, err
	}

	result := &ormDelegate{db}
	return result, nil
}

func (orm *ormDelegate) IsConnected() bool {
	err := orm.Ping()
	return err == nil
}

func (orm *ormDelegate) SaveContact(c *Contact) *Contact       { return c }
func (orm *ormDelegate) SaveMemo(memo *Memo) *Memo             { return memo }
func (orm *ormDelegate) SaveUser(user *User) *User             { return user }
func (orm *ormDelegate) SaveSession(session *Session) *Session { return session }

func (orm *ormDelegate) DeleteContact(contact *Contact) error { return nil }
func (orm *ormDelegate) DeleteMemo(memo *Memo) error          { return nil }
func (orm *ormDelegate) DeleteUser(user *User) error          { return nil }
func (orm *ormDelegate) DeleteSession(session *Session) error { return nil }
