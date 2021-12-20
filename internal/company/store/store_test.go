package companystore_test

import (
	companystore "github.com/emdeha/screener-go/internal/company/store"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompanyStore", func() {
	var (
		store *companystore.Store
		err   error
	)

	BeforeEach(func() {
		store = companystore.New(db)
	})

	When("InsertDocument", func() {
		var (
			doc string
		)

		JustBeforeEach(func() {
			err = store.InsertDocument(doc)
		})

		It("doesn't return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
