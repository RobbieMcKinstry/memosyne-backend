package spec

import (
	. "../model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"time"
)

const (
	PHONY_DB string = "phony.db"
)

func TestORM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Object relational mapper")
}

var _ = Describe("ORM", func() {
	Context("When I try to make a database", func() {
		It("should not return an error", func() {
			_, err := NewORM(PHONY_DB)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should be connected", func() {
			orm, err := NewORM(PHONY_DB)
			Expect(err).NotTo(HaveOccurred())
			Expect(orm.IsConnected()).To(Equal(true))
		})
	})
	Context("If I'm working with the database", func() {
		orm, _ := NewORM(PHONY_DB)

		// TODO add a test to verify that updating an object works
		Context("and I try and make a new user object", func() {

			user := &User{
				PhoneNum:  "412-445-3171",
				Email:      "thesnowmancometh@gmail.com",
				FirstName: "Robbie",
				LastName:  "McKinstry",
				Password:   "foobar",
			}
			orm.SaveUser(user)
			It("should have an ID", func() {
				Expect(user.UserId).NotTo(Equal(0))
			})
		})

		Context("and I try to make a new session object", func() {
			session := &Session{
				Expiration: time.Now().UTC().Format(time.RubyDate),
				UserId:    1,
			}
			orm.SaveSession(session)
			It("should have an ID", func() {
				Expect(session.SessionId).NotTo(Equal(0))
			})
		})

		Context("and I try to make a new memo object", func() {
			memo := &Memo{
				SenderId: 1,
				RecipientId: 2,
				Body: "Jenny please! I love you!",
				Time: time.Now().UTC().Format(time.RubyDate),
			}
			orm.SaveMemo(memo)
			PIt("should have an ID", func() {
				// Need to include a memo id to be able to delete individual memos
			})
		})

		Context("and I try to make a new contact object", func() {
			contact := &Contact{
				PhoneNum: "412-445-3171",
				Status: 2,
				FirstName: "Robbie",
				LastName: "McKinstry",
			}
			orm.SaveContact(contact)
			It("should have an ID", func() {
        Expect(contact.ContactId).NotTo(BeZero())
			})
		})

	})
})
