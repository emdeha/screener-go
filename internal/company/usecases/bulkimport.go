package usecases

import "context"

//go:generate counterfeiter . CompanyImporter
type CompanyImporter interface {
	DoImport(_ context.Context) error
}

type BulkImport struct {
	importer CompanyImporter
}

func NewBulkImport(importer CompanyImporter) *BulkImport {
	return &BulkImport{
		importer: importer,
	}
}

func (m *BulkImport) Do(ctx context.Context) error {
	return m.importer.DoImport(ctx)
}
