package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/jllanes-ss/avisos/account/domain"
	"github.com/jllanes-ss/avisos/account/domain/apperrors"

	"github.com/google/uuid"
)

type postgressUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &postgressUserRepository{
		DB: db,
	}
}

// FindByID fetches user by id
func (pg *postgressUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	query := `SELECT uid,name,email  FROM users WHERE uid = $1`

	err := pg.DB.QueryRowContext(ctx, query, uid).Scan(&user.ID, &user.Name, &user.Email)

	switch {
	case err == sql.ErrNoRows:
		return user, apperrors.NewNotFound("uid", uid.String())
	case err != nil:
		return user, err
	default:
		return user, err
	}
}

// FindByEmail retrieves user row by email address
func (pg *postgressUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	query := `SELECT uid,name,email,password  FROM users WHERE email = $1`

	err := pg.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("Unable to get user with email address: %v. Err: %v\n", email, err)
		return user, apperrors.NewNotFound("email", email)
	case err != nil:
		return user, err
	default:
		return user, err
	}
}
