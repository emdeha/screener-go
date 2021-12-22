package company_test

import (
	"context"

	"github.com/emdeha/screener-go/internal/company"
	"github.com/emdeha/screener-go/internal/company/companyfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	var (
		manager *company.Manager
		store   *companyfakes.FakeCompanyStore
		err     error
		ctx     context.Context
	)

	BeforeEach(func() {
		store = &companyfakes.FakeCompanyStore{}
		manager = company.New(store)
		ctx = context.Background()
	})

	When("InsertCompany", func() {
		JustBeforeEach(func() {
			err = manager.InsertCompany(ctx, &company.Company{})
		})

		It("doesn't return an error", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(store.InsertCompanyCallCount()).To(Equal(1))
		})
	})
})
