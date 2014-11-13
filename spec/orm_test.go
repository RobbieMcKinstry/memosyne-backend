package spec

import (
	. "hackathon/model"

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
				Phone_num:  "412-445-3171",
				Email:      "thesnowmancometh@gmail.com",
				First_name: "Robbie",
				Last_name:  "McKinstry",
				Password:   "foobar",
			}
			user = orm.SaveUser(user)
			It("should have an ID", func() {
				Expect(user.User_id).NotTo(Equal(0))
			})
		})

		Context("and I try to make a new session object", func() {
			session := &Session{
				Expiration: time.Now().UTC().Format(time.RubyDate),
				User_id:    1,
			}
			session = orm.SaveSession(session)
			It("should have an ID", func() {
				Expect(session.Session_id).NotTo(Equal(0))
			})
		})

	})
})
