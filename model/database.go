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

func NewORM(connectionString string) (ORM, err) {

	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return nil, err
	}

	result := &ormImplementation{ db }
	return result
}

func (orm ORM) SaveContact(c *Contact) *Contact {}
func (orm ORM) SaveMemo(memo *Memo) *Memo {}
func (orm ORM) SaveUser(user *User) *User {}
func (orm ORM) SaveSession(session *Session) *Session {}

func (orm ORM) DeleteContact(contact *Contact) error {}
func (orm ORM) DeleteMemo(memo *Memo) error {}
func (orm ORM) DeleteUser(user *User) error {}
func (orm ORM) DeleteSession(session *Session) error {}

