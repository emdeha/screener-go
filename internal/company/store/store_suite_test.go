package store_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Store Suite")
}

var (
	db     *mongo.Client
	dbName string
)

var _ = BeforeSuite(func() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	Expect(err).ToNot(HaveOccurred())

	dbName = "screener-test"
})

var _ = AfterSuite(func() {
	err := db.Database(dbName).Drop(context.Background())
	Expect(err).ToNot(HaveOccurred())
	Expect(db.Disconnect(context.Background()))
})
