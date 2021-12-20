package importer_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/emdeha/screener-go/internal/company"
	"github.com/emdeha/screener-go/internal/company/companyfakes"
	"github.com/emdeha/screener-go/internal/company/importer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EDGAR", func() {
	var (
		manager       *company.Manager
		companyStore  *companyfakes.FakeCompanyStore
		edgarImporter *importer.EDGAR
		err           error
		ctx           context.Context
	)

	BeforeEach(func() {
		companyStore = &companyfakes.FakeCompanyStore{}
		manager = company.New(companyStore, edgarImporter)
		edgarImporter = importer.New(manager)
	})

	When("ImportFile", func() {
		var (
			file        io.Reader
			companyData []byte
		)

		JustBeforeEach(func() {
			file = bytes.NewReader(companyData)
			err = edgarImporter.ImportFile(ctx, file)
		})

		Context("with proper json", func() {
			BeforeEach(func() {
				companyData, err = json.Marshal(company.Company{
					CIK:        1750,
					EntityName: "AAR CORP.",
				})
				Expect(err).ToNot(HaveOccurred())
			})

			It("imports file successfully", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(companyStore.InsertCompanyCallCount()).To(Equal(1))
			})
		})
		Context("with invalid json", func() {
			BeforeEach(func() {
				companyData = []byte("{ not: 'json")
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
				_, ok := err.(*json.SyntaxError)
				Expect(ok).To(BeTrue())
			})
		})
	})
})
