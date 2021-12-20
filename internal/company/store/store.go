package store

import (
	"context"

	"github.com/emdeha/screener-go/internal/company"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db           *mongo.Client
	databaseName string
}

func New(db *mongo.Client, databaseName string) *Store {
	return &Store{
		db:           db,
		databaseName: databaseName,
	}
}

func (s *Store) InsertCompany(
	ctx context.Context,
	companyData *company.Company,
) error {
	collection := s.db.Database(s.databaseName).Collection("companies")
	var (
		err error
	)
	_, err = collection.InsertOne(ctx, companyData)
	return err
}
