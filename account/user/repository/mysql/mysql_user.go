package mysql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jllanes-ss/avisos/account/domain"
	"github.com/jllanes-ss/avisos/account/domain/apperrors"
	//"github.com/jllanes-ss/avisos/account/user/repository/mysql"
)

type mysqlUserRepository struct {
	DB *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (m *mysqlUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	query := `SELECT id,name, email FROM tbl_user WHERE ID = ?`

	err := m.DB.QueryRowContext(ctx, query, uid.String()).Scan(&user.ID, &user.Name, &user.Email)

	switch {
	case err == sql.ErrNoRows:
		return user, apperrors.NewNotFound("uid", uid.String())
	case err != nil: //@todo need to actually check errors as it could be something other than not found
		return user, err
	default:
		return user, err
	}
}

func (m *mysqlUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	query := `SELECT id, name, email, password  FROM tbl_user WHERE email = ?`

	err := m.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	switch {
	case err == sql.ErrNoRows:
		return user, apperrors.NewNotFound("email", email)
	case err != nil: //@todo need to actually check errors as it could be something other than not found
		return user, err
	default:
		return user, err
	}
}
