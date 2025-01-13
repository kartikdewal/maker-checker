package psql

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"maker-checker/logger"
	"net/url"
	"strings"
)

type Config struct {
	User               string
	Password           string
	Host               string
	Port               string
	DbName             string
	ServiceName        string
	SkipMigrations     bool
	MigrationsLocation string
}

func NewConnection(cfg *Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.DbName, cfg.Host, cfg.Port)
	db, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(0)
	return db, nil
}

// RunMigrations runs the up migrations from the `/store/psql/migrations` directory on application start.
// If no new migrations are found, it logs a message and returns.
func RunMigrations(ctx context.Context, log logger.ContextLogger, config *Config) {
	if config.SkipMigrations {
		return
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(config.User),
		url.QueryEscape(config.Password),
		url.QueryEscape(config.Host),
		config.Port,
		url.QueryEscape(config.DbName),
	)

	m, err := migrate.New(config.MigrationsLocation, connStr)

	if err != nil {
		var errString = err.Error()
		for _, s := range []string{config.User, config.Password, config.Host, config.DbName} {
			// replace sensitive data with ***
			errString = strings.ReplaceAll(errString, s, "***")
			errString = strings.ReplaceAll(errString, url.QueryEscape(s), "***")
		}
		log.Fatalf(ctx, "failed to run DB migrations: %+v", errString)
		return
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info(ctx, "no migrations need running")
		} else {
			log.Fatalf(ctx, "failed to run DB migrations: %+v", err)
		}
		return
	}

	log.Info(ctx, "DB migrations ran successfully")
}
