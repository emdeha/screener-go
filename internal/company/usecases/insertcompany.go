package usecases

import (
	"context"

	"github.com/emdeha/screener-go/internal/company"
)

//go:generate counterfeiter . CompanyStore
type CompanyStore interface {
	InsertCompany(_ context.Context, _ *company.Company) error
}

type InsertCompany struct {
	store CompanyStore
}

func NewInsertCompany(store CompanyStore) *InsertCompany {
	return &InsertCompany{
		store: store,
	}
}

func (m *InsertCompany) Do(ctx context.Context, data *company.Company) error {
	return m.store.InsertCompany(ctx, data)
}
