package usecases_test

import (
	"context"

	"github.com/emdeha/screener-go/internal/company"
	"github.com/emdeha/screener-go/internal/company/usecases"
	"github.com/emdeha/screener-go/internal/company/usecases/usecasesfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InsertCompany", func() {
	var (
		insertCompany *usecases.InsertCompany
		store         *usecasesfakes.FakeCompanyStore
		err           error
		ctx           context.Context
	)

	BeforeEach(func() {
		store = &usecasesfakes.FakeCompanyStore{}
		insertCompany = usecases.NewInsertCompany(store)
		ctx = context.Background()
	})

	When("InsertCompany", func() {
		JustBeforeEach(func() {
			err = insertCompany.Do(ctx, &company.Company{})
		})

		It("doesn't return an error", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(store.InsertCompanyCallCount()).To(Equal(1))
		})
	})
})
