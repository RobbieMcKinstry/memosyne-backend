package spec

import (
	. "hackathon/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const (
	PHONY_DB string = "phony.db"
)

func TestBooks(t *testing.T) {
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
			orm, _ := NewORM(PHONY_DB)
			Expect(orm.IsConnected()).To(Equal(true))
		})
	})
/*	Context("If I'm working with the database", func() {
		Context("and I try and make a new user object", func() {
			orm, _ := NewORM(PHONY_DB)
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
	})
*/
})
