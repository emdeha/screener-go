package companystore_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCompanyStore(t *testing.T) {
	RegisterFailHandler(Fail)
}

var (
	db *mongo.Client
)

var _ = BeforeSuite(func() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://local:27017"))
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(db.Disconnect(context.Background()))
})
