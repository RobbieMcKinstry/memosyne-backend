package dummy

import (
    "database/sql"
    "fmt"
    _ "github.com/mxk/go-sqlite/sqlite3"
    "time"
)

func main() {

    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
    }

    db.Exec("PRAGMA foreign_keys = ON;")

    cid := 3
    phone_num := "412-444-4444"
    status := 1
    db.Exec("INSERT INTO Contact VALUES (?, ?, ?)", cid, phone_num, status)

    userPhoneNumber := "111-222-3333"
    userFirstName := "Bill"
    userLastName := "Blah"
    userId := 1
    contactRef := 3    
    db.Exec("INSERT INTO User VALUES (\"111-222-3333\", \"Bill\", \"Mike\", 1, 3, \"password\")")

    timeVal := time.Now()
    db.Exec("INSERT INTO Memo VALUES (1, 2, \"Lots of information\", " + timeVal + ")")

    db.Exec("INSERT INTO Contact_Reference VALUES (8, 9)")

    db.Exec("INSERT INTO Session VALUES (109, " + time.Now() + ", 20)")

    db.Close()
}