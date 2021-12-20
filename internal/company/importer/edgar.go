package importer

import (
	"context"
	"encoding/json"
	"io"

	"github.com/emdeha/screener-go/internal/company"
)

type EDGAR struct {
	companyManager *company.Manager
}

func New(companyManager *company.Manager) *EDGAR {
	return &EDGAR{
		companyManager: companyManager,
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

func (e *EDGAR) DoImport() error {
	return nil
}
