package main

func (this *Contact) newContact() *Contact {
    ret := new (Contact)
    
    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
    }

    rows, err := db.Query("SELECT * FROM Contact WHERE Contact.name=?", name)
    if err != nil {
        fmt.Println(err)
    }

    if rows != nil {
        /* user exists! */

    }

    db.Close()

    return ret
}

func (this *Contact) delete() bool {

    ret := false

    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
    }

    result, err := db.Exec("DELETE FROM Contact WHERE Contact.id=?", id)
    if err != nil {
        fmt.Println(err)
        ret = true
    }

    db.Close()
    return ret
}

func (this *Contact) isAccepted() bool {



    return true
}

func (this *Contact) save() {

}

