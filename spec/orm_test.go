package spec

import (
	. "hackathon/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)
func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite") 
}

var _ = Describe("ORM", func() {
	Context("When I try to make a database", func() {
		It("should not return an error", func() {
			_, err := NewORM("phony.db")
			Expect(err).NotTo(HaveOccurred())
		})
		It("should be connected", func() {
			orm, _ := NewORM("phony.db")
			Expect(orm.IsConnected()).To(Equal(true))
		})
	})
})
