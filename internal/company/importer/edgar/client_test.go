package importer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/emdeha/screener-go/internal/company"
	importer "github.com/emdeha/screener-go/internal/company/importer/edgar"
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func zipResponse(resp http.ResponseWriter) (*bytes.Buffer, error) {
	firstCompany := company.Company{
		CIK:        1234567,
		EntityName: "First",
	}
	fc, err := json.Marshal(firstCompany)
	if err != nil {
		return nil, err
	}

	secondCompany := company.Company{
		CIK:        7654321,
		EntityName: "Second",
	}
	sc, err := json.Marshal(secondCompany)
	if err != nil {
		return nil, err
	}

	archiveContents := []archive{
		{"CIK0001234567.json", string(fc)},
		{"CIK0007654321.json", string(sc)},
	}

	archiveResponse, err := writeToArchive(archiveContents)
	if err != nil {
		return nil, err
	}

	resp.Header().Set("content-type", "application/zip")
	_, err = resp.Write(archiveResponse.Bytes())
	if err != nil {
		return nil, err
	}

	return archiveResponse, nil
}

var _ = Describe("EDGARClient", func() {
	When("GetBulkData", func() {
		var (
			data        []byte
			err         error
			handler     http.Handler
			server      *httptest.Server
			edgarClient *importer.Client
		)

		JustBeforeEach(func() {
			server = httptest.NewServer(handler)

			var parsedURL *url.URL
			parsedURL, err = url.Parse(server.URL)
			Expect(err).ToNot(HaveOccurred())
			edgarClient = importer.NewEDGARClient(parsedURL, "test@test.com")
			data, err = edgarClient.GetBulkData(context.Background())
		})
		JustAfterEach(func() {
			server.Close()
		})

		Context("no error during retrieval", func() {
			var (
				archiveResponse *bytes.Buffer
			)

			BeforeEach(func() {
				handler = http.HandlerFunc(
					func(resp http.ResponseWriter, req *http.Request) {
						var zipErr error
						archiveResponse, zipErr = zipResponse(resp)
						Expect(zipErr).ToNot(HaveOccurred())
					},
				)
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(Equal(archiveResponse.Bytes()))
			})
		})

		Context("with error from the endpoint", func() {
			BeforeEach(func() {
				handler = http.HandlerFunc(
					func(resp http.ResponseWriter, req *http.Request) {
						http.Error(resp, "test", http.StatusUnauthorized)
					},
				)
			})

			It("fails with 403", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("401"))
			})
		})
	})
})
