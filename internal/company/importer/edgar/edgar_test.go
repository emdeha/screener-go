package edgarimporter_test

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/emdeha/screener-go/internal/company"
	"github.com/emdeha/screener-go/internal/company/companyfakes"
	edgarimporter "github.com/emdeha/screener-go/internal/company/importer/edgar"
	"github.com/emdeha/screener-go/internal/company/importer/edgar/edgarfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type archive struct {
	Name, Body string
}

func writeToArchive(archiveContents []archive) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	archiveWriter := zip.NewWriter(buf)
	for _, file := range archiveContents {
		f, err := archiveWriter.Create(file.Name)
		if err != nil {
			return nil, err
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			return nil, err
		}
	}
	err := archiveWriter.Close()
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// TODO: Something eats hard-drive when one of the tests fails. About 20Mb per
// run. Should see why is that.
// To reclaim space, just clean up /tmp/ginkgo*.
var _ = Describe("EDGAR", func() {
	var (
		manager       *company.Manager
		companyStore  *companyfakes.FakeCompanyStore
		edgarClient   *edgarfakes.FakeEDGARClient
		edgarImporter *edgarimporter.EDGAR
		err           error
		ctx           context.Context
	)

	BeforeEach(func() {
		companyStore = &companyfakes.FakeCompanyStore{}
		manager = company.New(companyStore)
		edgarClient = &edgarfakes.FakeEDGARClient{}
		edgarImporter = edgarimporter.New(manager, edgarClient)
	})

	When("ImportFile", func() {
		var (
			companyData []byte
		)

		JustBeforeEach(func() {
			file := bytes.NewReader(companyData)
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

	When("ImportFilesFromArchive", func() {
		var (
			archiveContents []archive

			firstCompany, secondCompany company.Company
		)

		JustBeforeEach(func() {
			var buf *bytes.Buffer
			buf, err = writeToArchive(archiveContents)
			Expect(err).ToNot(HaveOccurred())

			archiveReader, newReaderErr := zip.NewReader(
				bytes.NewReader(buf.Bytes()), int64(len(buf.Bytes())))
			Expect(newReaderErr).ToNot(HaveOccurred())
			err = edgarImporter.ImportFilesFromArchive(ctx, archiveReader)
		})

		Context("with proper archive", func() {
			BeforeEach(func() {
				firstCompany = company.Company{
					CIK:        1234567,
					EntityName: "First",
				}
				fc, err := json.Marshal(firstCompany)
				Expect(err).ToNot(HaveOccurred())

				secondCompany = company.Company{
					CIK:        7654321,
					EntityName: "Second",
				}
				sc, err := json.Marshal(secondCompany)
				Expect(err).ToNot(HaveOccurred())

				archiveContents = []archive{
					{"CIK0001234567.json", string(fc)},
					{"CIK0007654321.json", string(sc)},
				}
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(companyStore.InsertCompanyCallCount()).To(Equal(2))
				_, fc := companyStore.InsertCompanyArgsForCall(0)
				Expect(*fc).To(Equal(firstCompany))
				_, sc := companyStore.InsertCompanyArgsForCall(1)
				Expect(*sc).To(Equal(secondCompany))
			})
		})
		Context("with no companies in the archive", func() {
			BeforeEach(func() {
				archiveContents = []archive{}
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(companyStore.InsertCompanyCallCount()).To(Equal(0))
			})
		})
		Context("with invalid json in archive", func() {
			BeforeEach(func() {
				archiveContents = []archive{
					{"CIK0001234567", "invalid"},
				}
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	When("DoImport", func() {
		var getBulkDataError error

		JustBeforeEach(func() {
			var firstCompany []byte
			firstCompany, err = json.Marshal(company.Company{
				CIK:        1234567,
				EntityName: "First",
			})
			Expect(err).ToNot(HaveOccurred())

			var secondCompany []byte
			secondCompany, err = json.Marshal(company.Company{
				CIK:        7654321,
				EntityName: "Second",
			})
			Expect(err).ToNot(HaveOccurred())

			archiveContents := []archive{
				{"CIK0001234567.json", string(firstCompany)},
				{"CIK0007654321.json", string(secondCompany)},
			}

			var buf *bytes.Buffer
			buf, err = writeToArchive(archiveContents)
			Expect(err).ToNot(HaveOccurred())

			edgarClient.GetBulkDataReturns(buf.Bytes(), getBulkDataError)
			err = edgarImporter.DoImport(ctx)
		})

		Context("no error from GetBulkData", func() {
			BeforeEach(func() {
				getBulkDataError = nil
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(companyStore.InsertCompanyCallCount()).To(Equal(2))
			})
		})

		Context("error from GetBulkData", func() {
			BeforeEach(func() {
				getBulkDataError = errors.New("test")
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
