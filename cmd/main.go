package main

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caarlos0/env"
	edgarimporter "github.com/emdeha/screener-go/internal/company/importer/edgar"
	companystore "github.com/emdeha/screener-go/internal/company/store"
	companyusecases "github.com/emdeha/screener-go/internal/company/usecases"
)

type Config struct {
	DBUrl  string `env:"DB_URL" envDefault:"mongodb://127.0.0.1:27017"`
	DBName string `env:"DB_NAME" envDefault:"scanner"`

	EDGARUrl       string `env:"EDGAR_URL" envDefault:"https://www.sec.gov/Archives/edgar/daily-index/xbrl/companyfacts.zip"`
	EDGARUserAgent string `env:"EDGAR_USER_AGENT" envDefault:"test@test.com"`
}

func main() {
	ctx := context.Background()
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	insertCompany := setupInsertCompany(ctx, cfg)
	bulkImport := setupBulkImport(ctx, cfg, insertCompany)

	err = bulkImport.Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func setupInsertCompany(
	ctx context.Context, cfg *Config,
) *companyusecases.InsertCompany {
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DBUrl))
	if err != nil {
		log.Fatal(err)
	}
	companyStore := companystore.New(db, cfg.DBName)

	return companyusecases.NewInsertCompany(companyStore)
}

func setupBulkImport(
	ctx context.Context,
	cfg *Config,
	insertCompany *companyusecases.InsertCompany,
) *companyusecases.BulkImport {
	edgarEndpoint, err := url.Parse(cfg.EDGARUrl)
	if err != nil {
		log.Fatal(err)
	}
	edgarClient := edgarimporter.NewEDGARClient(
		edgarEndpoint, cfg.EDGARUserAgent)
	companyImporter := edgarimporter.New(insertCompany, edgarClient)

	return companyusecases.NewBulkImport(companyImporter)
}
