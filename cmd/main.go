package main

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/emdeha/screener-go/internal/company"
	importer "github.com/emdeha/screener-go/internal/company/importer/edgar"
	"github.com/emdeha/screener-go/internal/company/store"
)

func main() {
	ctx := context.Background()

	companyManager := setupCompanyManager(ctx)
	companyImporter := setupCompanyImporter(ctx, companyManager)

	err := companyImporter.DoImport(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func setupCompanyManager(ctx context.Context) *company.Manager {
	dbURL := "mongodb://127.0.0.1:27017"
	dbName := "scanner"
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	companyStore := store.New(db, dbName)

	return company.New(companyStore)
}

func setupCompanyImporter(
	ctx context.Context,
	companyManager *company.Manager,
) *importer.EDGAR {
	edgarEndpoint, err := url.Parse(
		"https://www.sec.gov/Archives/edgar/daily-index/xbrl/companyfacts.zip")
	if err != nil {
		log.Fatal(err)
	}
	userAgent := "cpi.cecko@gmail.com"
	edgarClient := importer.NewEDGARClient(edgarEndpoint, userAgent)
	return importer.New(companyManager, edgarClient)
}
