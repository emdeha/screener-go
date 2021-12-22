package usecases_test

import (
	"context"

	"github.com/emdeha/screener-go/internal/company/usecases"
	"github.com/emdeha/screener-go/internal/company/usecases/usecasesfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BulkImport", func() {
	var (
		bulkImport *usecases.BulkImport
		importer   *usecasesfakes.FakeCompanyImporter
		err        error
		ctx        context.Context
	)

	BeforeEach(func() {
		importer = &usecasesfakes.FakeCompanyImporter{}
		bulkImport = usecases.NewBulkImport(importer)
		ctx = context.Background()
	})

	When("BulkImport", func() {
		JustBeforeEach(func() {
			err = bulkImport.Do(ctx)
		})

		It("doesn't return an error", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(importer.DoImportCallCount()).To(Equal(1))
		})
	})
})
