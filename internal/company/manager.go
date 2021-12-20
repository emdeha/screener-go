package company

import "context"

type Company struct {
	CIK        int    `bson:"cik"`
	EntityName string `bson:"entityName"`
	Facts      Facts  `bson:"facts"`
}

type Facts struct {
	DEI    interface{} `bson:"dei"`
	USGaap interface{} `bson:"us-gaap"`
}

//go:generate counterfeiter . Importer
type Importer interface {
	DoImport() error
}

//go:generate counterfeiter . CompanyStore
type CompanyStore interface {
	InsertCompany(_ context.Context, _ *Company) error
}

type Manager struct {
	store    CompanyStore
	importer Importer
}

func New(store CompanyStore, importer Importer) *Manager {
	return &Manager{
		store:    store,
		importer: importer,
	}
}

func (m *Manager) InsertCompany(ctx context.Context, data *Company) error {
	return m.store.InsertCompany(ctx, data)
}
