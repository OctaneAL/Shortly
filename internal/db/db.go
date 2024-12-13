package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(connectionString string) *DB {
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	return &DB{Pool: pool}
}

func (db *DB) Close() {
	db.Pool.Close()
}

type URLMapping struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}

func (db *DB) SaveURL(ctx context.Context, originalURL, shortCode string) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO url_mappings (original_url, short_code) 
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, originalURL, shortCode)
	return err
}

func (db *DB) GetOriginalURL(ctx context.Context, shortCode string) (*URLMapping, error) {
	row := db.Pool.QueryRow(ctx, `
		SELECT id, original_url, short_code, created_at 
		FROM url_mappings 
		WHERE short_code = $1
	`, shortCode)

	var mapping URLMapping
	err := row.Scan(&mapping.ID, &mapping.OriginalURL, &mapping.ShortCode, &mapping.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &mapping, nil
}
