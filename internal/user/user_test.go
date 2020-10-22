package user

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"simpl-coding-challenge/internal/infrastructure"
)

var _ = Describe("Test User", func() {
	Context("For on-boarding new user flow", func() {
		It("In case if username is unique", func() {
			userClient := NewUser(infrastructure.NewMockMySQL())
			message, err := userClient.AddNewUser("mourya", "mourya.g9@gmail.com", 1000)
			Expect(err).To(BeNil())
			Expect(message).To(Equal("mourya(1000)"))
		})
		It("In case if username is unique", func() {
			userClient := NewUser(infrastructure.NewMockMySQL().WithInsertError(errors.New("can't create a duplicate name")))
			message, err := userClient.AddNewUser("mourya", "mourya.g9@gmail.com", 1000)
			Expect(err.Error()).To(Equal("can't create a duplicate name"))
			Expect(message).To(Equal(""))
		})
	})

	Context("For testing user dues flow", func() {
		It("When no users are present", func() {
			userClient := NewUser(infrastructure.NewMockMySQL().WithQueryRows(infrastructure.NewMockRows().WithRecordCount(0)))
			dueAmount, err := userClient.GetTotalUserDues()
			Expect(err).To(BeNil())
			Expect(dueAmount).To(Equal(int64(0)))
		})
	})
})
