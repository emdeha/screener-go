package companystore

import "go.mongodb.org/mongo-driver/mongo"

type Store struct {
	db *mongo.Client
}

func New(db *mongo.Client) *Store {
	return &Store{db: db}
}

func (s *Store) InsertDocument(_ string) error {
	return nil
}
