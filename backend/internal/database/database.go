package database

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed tables.sql
var Tables string

var dbpool *pgxpool.Pool

func Start() {
	connect()
	initialize()
}

func Stop() {
	dbpool.Close()
}

func connect() {
	db, err := pgxpool.NewWithConfig(context.Background(), config())
	if err != nil {
		fmt.Println("Error while creating connection to the database!!")
	}

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database!!")
	dbpool = db
}

func config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5


	dbConfig, err := pgxpool.ParseConfig(os.Getenv("DBURI"))
	if err != nil {
		fmt.Println("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
	}

	return dbConfig
}

func initialize() {
	_, err := dbpool.Exec(context.Background(), Tables)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create tables: %v\n", err)
		os.Exit(1)
	}
}
