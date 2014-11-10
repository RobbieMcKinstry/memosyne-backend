package model

import (
	"database/sql"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

type ORM interface {
	SaveContact(*Contact) *Contact
	SaveMemo(*Memo)       *Memo
	SaveUser(*User)	      *User
	SaveSession(*Session) *Session

	DeleteContact(*Contact) error
	DeleteMemo(*Memo)       error
	DeleteUser(*User)       error
	DeleteSession(*Session) error
}

type ormImplementation struct {
	*sql.DB
}

func NewORM(connectionString string) (ORM, error) {

	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return nil, err
	}

	result := &ormImplementation{ db }
	return result, nil
}

func (orm *ormImplementation) SaveContact(c *Contact) *Contact { return c}
func (orm *ormImplementation) SaveMemo(memo *Memo) *Memo {return memo}
func (orm *ormImplementation) SaveUser(user *User) *User {return user}
func (orm *ormImplementation) SaveSession(session *Session) *Session {return session}

func (orm *ormImplementation) DeleteContact(contact *Contact) error {return nil}
func (orm *ormImplementation) DeleteMemo(memo *Memo) error {return nil}
func (orm *ormImplementation) DeleteUser(user *User) error {return nil}
func (orm *ormImplementation) DeleteSession(session *Session) error {return nil}

