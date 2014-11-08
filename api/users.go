package api

type User struct {
	phone_num   string
	first_name  string
	last_name   string
	user_id     int
	contact_ref int
	password    string
}

func (this *User) UserSave() {

}

func (this *User) UserNew() *User {
	var newUser User
	return *newUser
}

func (this *User) UserDelete() bool {
	return true
}

func (this *User) GetContacts() []*Contact {
	var contactSlice []Contact
	return *contactSlice
}

func (this *User) GetMemos() []*Memos {
	var memoSlice []Memo
	return *memoSlice
}
