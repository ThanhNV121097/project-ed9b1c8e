package db

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.up.sql
var migrationFiles embed.FS

var ErrNoRows = pgx.ErrNoRows

func Migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(ctx, `create table if not exists schema_migrations (
		version text primary key,
		applied_at timestamptz not null default now()
	)`); err != nil {
		return err
	}

	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return err
	}

	var names []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".up.sql") {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)

	for _, name := range names {
		if err := apply(ctx, pool, name); err != nil {
			return fmt.Errorf("apply %s: %w", name, err)
		}
	}
	return nil
}

func apply(ctx context.Context, pool *pgxpool.Pool, name string) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var exists bool
	err = tx.QueryRow(ctx, "select exists(select 1 from schema_migrations where version = $1)", name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return tx.Commit(ctx)
	}

	sqlBytes, err := migrationFiles.ReadFile("migrations/" + name)
	if err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, "insert into schema_migrations (version) values ($1)", name); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func IsNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
