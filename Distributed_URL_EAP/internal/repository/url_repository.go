// package repository

// import (
// 	"context"
// 	"errors"

// 	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// var ErrURLNotFound = errors.New("url not found")

// type UrlRepository interface {
// 	Create(ctx context.Context, url *model.UrlDB) error
// 	GetbyShortCode(ctx context.Context, shortCode string) (*model.UrlDB, error)
// 	DeleteByShortCode(ctx context.Context, shortCode string) error
// 	IncrementClicked(ctx context.Context, shortCode string) error
// }

// type PostgresUrlRepository struct {
// 	db *pgxpool.Pool
// }

// // Create implements [UrlRepository].
// func (p *PostgresUrlRepository) Create(ctx context.Context, url *model.UrlDB) error {
// 	query := `INSERT INTO urls (id, url, clicked, short_code, created_at) VALUE ($1, $2, $3, $4, $5)`

// 	_, err := p.db.Exec(ctx, query, url.Id, url.Url, url.Clicked, url.ShortCode, url.CreatedAt)
// 	return err
// }

// // DeleteByShortCode implements [UrlRepository].
// func (p *PostgresUrlRepository) DeleteByShortCode(ctx context.Context, shortCode string) error {

// }

// // GetbyShortCode implements [UrlRepository].
// func (p *PostgresUrlRepository) GetbyShortCode(ctx context.Context, shortCode string) (*model.UrlDB, error) {
// 	query := `SELECT id, url, clicked, short_code, created_at FROM urls WHERE short_code = $1`

// 	var url model.UrlDB
// 	err := p.db.QueryRow(ctx, query, shortCode).Scan(
// 		&url.Id,
// 		&url.Url,
// 		&url.Clicked,
// 		&url.ShortCode,
// 		&url.CreatedAt,
// 	)

// 	if errors.Is(err, pgx.ErrNoRows) {
// 		return nil, ErrURLNotFound
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &url, nil
// }

// // IncrementClicked implements [UrlRepository].
// func (p *PostgresUrlRepository) IncrementClicked(ctx context.Context, shortCode string) error {
// 	panic("unimplemented")
// }

// func NewPostgresUrlRepository(db *pgxpool.Pool) *PostgresUrlRepository {
// 	return &PostgresUrlRepository{db: db}
// }

package repository

import (
	"context"
	"errors"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrURLNotFound = errors.New("url not found")

type UrlRepository interface {
	Create(ctx context.Context, url *model.UrlDB) error
	GetByShortCode(ctx context.Context, shortCode string) (*model.UrlDB, error)
	DeleteByShortCode(ctx context.Context, shortCode string) error
	IncrementClicked(ctx context.Context, shortCode string) error
}

type PostgresUrlRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUrlRepository(db *pgxpool.Pool) *PostgresUrlRepository {
	return &PostgresUrlRepository{
		db: db,
	}
}

func (r *PostgresUrlRepository) Create(
	ctx context.Context,
	url *model.UrlDB,
) error {

	query := `
		INSERT INTO urls (
			id,
			url,
			clicked,
			short_code,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		url.Id,
		url.Url,
		url.Clicked,
		url.ShortCode,
		url.CreatedAt,
	)

	return err
}

func (r *PostgresUrlRepository) GetByShortCode(
	ctx context.Context,
	shortCode string,
) (*model.UrlDB, error) {

	query := `
		SELECT
			id,
			url,
			clicked,
			short_code,
			created_at
		FROM urls
		WHERE short_code = $1
	`

	var url model.UrlDB

	err := r.db.QueryRow(
		ctx,
		query,
		shortCode,
	).Scan(
		&url.Id,
		&url.Url,
		&url.Clicked,
		&url.ShortCode,
		&url.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrURLNotFound
	}

	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *PostgresUrlRepository) DeleteByShortCode(
	ctx context.Context,
	shortCode string,
) error {

	query := `
		DELETE FROM urls
		WHERE short_code = $1
	`

	result, err := r.db.Exec(
		ctx,
		query,
		shortCode,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrURLNotFound
	}

	return nil
}

func (r *PostgresUrlRepository) IncrementClicked(
	ctx context.Context,
	shortCode string,
) error {

	query := `
		UPDATE urls
		SET clicked = clicked + 1
		WHERE short_code = $1
	`

	result, err := r.db.Exec(
		ctx,
		query,
		shortCode,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrURLNotFound
	}

	return nil
}
