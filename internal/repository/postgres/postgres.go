package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"practice3/pkg/modules"
)

type Dialect struct {
	DB *sqlx.DB
}

func New(cfg *modules.PostgreConfig) *Dialect {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password,
		cfg.DBName, cfg.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	runMigrations(cfg)

	return &Dialect{DB: db}
}

func runMigrations(cfg *modules.PostgreConfig) {
	sourceURL := "file://database/migrations"
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password,
		cfg.Host, cfg.Port,
		cfg.DBName, cfg.SSLMode,
	)

	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}