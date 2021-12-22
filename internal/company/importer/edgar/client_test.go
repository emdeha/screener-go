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
			edgarClient = importer.NewEDGARClient(parsedURL)
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
						firstCompany := company.Company{
							CIK:        1234567,
							EntityName: "First",
						}
						var fc []byte
						fc, err = json.Marshal(firstCompany)
						Expect(err).ToNot(HaveOccurred())

						secondCompany := company.Company{
							CIK:        7654321,
							EntityName: "Second",
						}
						var sc []byte
						sc, err = json.Marshal(secondCompany)
						Expect(err).ToNot(HaveOccurred())

						archiveContents := []archive{
							{"CIK0001234567.json", string(fc)},
							{"CIK0007654321.json", string(sc)},
						}

						archiveResponse, err = writeToArchive(archiveContents)
						Expect(err).ToNot(HaveOccurred())

						resp.Header().Set("content-type", "application/zip")
						_, err = resp.Write(archiveResponse.Bytes())
						Expect(err).ToNot(HaveOccurred())
					},
				)
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(Equal(archiveResponse.Bytes()))
			})
		})
	})
})
