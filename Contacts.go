package api

type Contact struct {
  cid int
  phone_number string
  status int
}

func (this *Contact) save() {

}

func (this *Contact) new() *Contact {
    ret := new (Contact)
    return ret
}

func (this *Contact) delete() bool {
    return true
}

func (this *Contact) isAccepted() bool {
    return true
}
