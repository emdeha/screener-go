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

	log.Println(cfg)

	insertCompany := setupInsertCompany(ctx, cfg)
	companyImporter := setupCompanyImporter(ctx, cfg, insertCompany)

	err = companyImporter.DoImport(ctx)
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

func setupCompanyImporter(
	ctx context.Context,
	cfg *Config,
	insertCompany *companyusecases.InsertCompany,
) *edgarimporter.EDGAR {
	edgarEndpoint, err := url.Parse(cfg.EDGARUrl)
	if err != nil {
		log.Fatal(err)
	}
	edgarClient := edgarimporter.NewEDGARClient(
		edgarEndpoint, cfg.EDGARUserAgent)
	return edgarimporter.New(insertCompany, edgarClient)
}
