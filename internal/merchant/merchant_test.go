package merchant

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"simpl-coding-challenge/internal/infrastructure"
)

var _ = Describe("Test User", func() {
	Context("For on-boarding new user flow", func() {
		It("In case if merchantName is unique", func() {
			merchantClient := NewMerchant(infrastructure.NewMockMySQL())
			message, err := merchantClient.AddNewMerchant("mourya", "mourya.g9@gmail.com", 10)
			Expect(err).To(BeNil())
			Expect(message).To(Equal("mourya(10.00)"))
		})
		It("In case if username is unique", func() {
			merchantClient := NewMerchant(infrastructure.NewMockMySQL().WithInsertError(errors.New("can't create a duplicate name")))
			message, err := merchantClient.AddNewMerchant("mourya", "mourya.g9@gmail.com", 10)
			Expect(err.Error()).To(Equal("can't create a duplicate name"))
			Expect(message).To(Equal(""))
		})
	})

	Context("For updating merchant discount flow", func() {
		It("In case of a successful flow", func() {
			merchantClient := NewMerchant(infrastructure.NewMockMySQL())
			message, err := merchantClient.UpdateDiscountPercentage("mourya", 20)
			Expect(err).To(BeNil())
			Expect(message).To(Equal("mourya(20.00)"))
		})
		It("In case if merchant name doesn't exist", func() {
			merchantClient := NewMerchant(infrastructure.NewMockMySQL().WithUpdateError(errors.New("internal Server Error")))
			message, err := merchantClient.UpdateDiscountPercentage("mourya", 10)
			Expect(err.Error()).To(Equal("internal Server Error"))
			Expect(message).To(Equal(""))
		})
	})
})
