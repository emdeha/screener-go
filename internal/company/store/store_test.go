// +build integration

package companystore_test

import (
	"context"

	"github.com/emdeha/screener-go/internal/company"
	companystore "github.com/emdeha/screener-go/internal/company/store"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Company Store", func() {
	var (
		companyStore *companystore.Store
		ctx          context.Context
		err          error
	)

	BeforeEach(func() {
		ctx = context.Background()
		companyStore = companystore.New(db, dbName)
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
