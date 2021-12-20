// +build integration

package store_test

import (
	"context"

	"github.com/emdeha/screener-go/internal/company"
	"github.com/emdeha/screener-go/internal/company/store"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	var (
		companyStore *store.Store
		ctx          context.Context
		err          error
	)

	BeforeEach(func() {
		ctx = context.Background()
		companyStore = store.New(db, "screener")
	})

	When("InsertCompany", func() {
		var (
			companyData company.Company
		)

		BeforeEach(func() {
			companyData = company.Company{
				CIK:        1750,
				EntityName: "AAR CORP.",
			}
		})
		JustBeforeEach(func() {
			err = companyStore.InsertCompany(ctx, &companyData)
		})

		It("doesn't return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
