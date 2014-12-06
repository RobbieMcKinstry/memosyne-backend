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

var _ = Describe("Utility Methods", func() {
	Context("When working with a session", func() {
		t := time.Now()
		session := &Session{
			SessionId:  1,
			Expiration: t,
			UserId:     1,
		}
		It("should be equal to another identity session.", func() {
			session2 := &Session{
				SessionId:  1,
				Expiration: t,
				UserId:     1,
			}
			Expect(session.Equals(session2)).To(BeTrue())
		})
	})
})

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
				Email:     "thesnowmancometh@gmail.com",
				FirstName: "Robbie",
				LastName:  "McKinstry",
				Password:  "foobar",
			}
			orm.SaveUser(user)
			It("should have an ID", func() {
				Expect(user.UserId).NotTo(Equal(0))
			})
		})

		Context("and I try to make a new session object", func() {
			session := &Session{
				Expiration: time.Now().Truncate(time.Second).UTC(),
				UserId:     1,
			}
			orm.SaveSession(session)
			It("should have an ID", func() {
				Expect(session.SessionId).NotTo(Equal(0))
			})

			It("should be able to be found by it's ID.", func() {
				id := session.SessionId
				sess := orm.FindSessionByID(id)
				Expect(session.Equals(sess)).To(BeTrue(), "\nSession 1: %v\nSession 2: %v", session.ToString(), sess.ToString())
			})
		})

		Context("and I try to make a new memo object", func() {
			memo := &Memo{
				SenderId:    1,
				RecipientId: 2,
				Body:        "Jenny please! I love you!",
				Time:        time.Now().Truncate(time.Second).UTC(),
			}
			orm.SaveMemo(memo)
			It("should have an ID", func() {
				Expect(memo.ID).NotTo(BeZero())
			})
		})

		Context("and I try to make a new contact object", func() {
			contact := &Contact{
				PhoneNum:  "412-445-3171",
				Status:    2,
				FirstName: "Robbie",
				LastName:  "McKinstry",
			}
			orm.SaveContact(contact)
			It("should have an ID", func() {
				Expect(contact.ContactId).NotTo(BeZero())
			})
		})

	})
})
