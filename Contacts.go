/* 
 * Methods for Contact class
 *      
 *      Author: Tyler Raborn
 */

package main

func (this *Contact) newContact() *Contact {
    ret := new (Contact)
    
    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
    }

    rows, err := db.Query("SELECT * FROM Contact WHERE Contact.cid=?", this.cid)
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

    result, err := db.Exec("DELETE FROM Contact WHERE Contact.cid=?", this.cid)
    if err != nil {
        fmt.Println(err)
        ret = true
    }

    db.Close()
    return ret
}

/* checks to see if the contact is accepted - unsure as to why this exists... */
/*
func (this *Contact) isAccepted(externalUser *User) bool {

    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
    }

    ret, err := db.Query("SELECT status FROM Contact WHERE cid=?", externalUser.cid)
    if err != nil {
        fmt.Println(err)
    }

    db.Close()
    return ret
}
*/

/* Saves contact data to db */
func (this *Contact) save() {

    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
    }

    db.Exec(
        "INSERT INTO Contact VALUES (?, ?, ?)", 
        this.cid, 
        this.phone_number, 
        this.status
    )
    
    db.Close()
}

