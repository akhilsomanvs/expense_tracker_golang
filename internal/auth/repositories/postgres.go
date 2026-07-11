package repositories

import (
	"context"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgressRepository(db *pgxpool.Pool) UserRepository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(
	ctx context.Context,
	user *models.UserModel,
) error {

	query := `
		INSERT INTO users
		(
			name,
			email,
			password_hash
		)
		VALUES
		($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *postgresRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.UserModel, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	user := &models.UserModel{}

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
