package importer

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/emdeha/screener-go/internal/company"
)

//go:generate counterfeiter . EDGARClient
type EDGARClient interface {
	// GetBulkData stores the whole file in memory as zip archives require to be
	// read in full in order to be unzipped.
	GetBulkData() ([]byte, error)
}

type EDGAR struct {
	companyManager *company.Manager
	edgarClient    EDGARClient
}

func New(companyManager *company.Manager, edgarClient EDGARClient) *EDGAR {
	return &EDGAR{
		companyManager: companyManager,
		edgarClient:    edgarClient,
	}
}

func (e *EDGAR) ImportFile(ctx context.Context, file io.Reader) error {
	var companyData company.Company

	// TODO: Add validation
	if err := json.NewDecoder(file).Decode(&companyData); err != nil {
		return err
	}
	return e.companyManager.InsertCompany(ctx, &companyData)
}

func (e *EDGAR) ImportFilesFromArchive(
	ctx context.Context,
	archive *zip.Reader,
) error {
	for _, file := range archive.File {
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		err = e.ImportFile(ctx, fileReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EDGAR) DoImport(ctx context.Context) error {
	bulkData, err := e.edgarClient.GetBulkData()
	if err != nil {
		return err
	}

	bulkDataAsArchive, err := zip.NewReader(
		bytes.NewReader(bulkData), int64(len(bulkData)))
	if err != nil {
		return err
	}

	return e.ImportFilesFromArchive(ctx, bulkDataAsArchive)
}
