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

//go:generate counterfeiter . CompanyStore
type CompanyStore interface {
	InsertCompany(_ context.Context, _ *Company) error
}

type Manager struct {
	store CompanyStore
}

func New(store CompanyStore) *Manager {
	return &Manager{
		store: store,
	}
}

func (m *Manager) InsertCompany(ctx context.Context, data *Company) error {
	return m.store.InsertCompany(ctx, data)
}
