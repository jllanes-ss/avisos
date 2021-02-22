package repository

import (
	"database/sql"
	"reflect"

	"github.com/getsentry/sentry-go"
	"github.com/jllanes-ss/avisos/account/domain"
	_userMysqlRepository "github.com/jllanes-ss/avisos/account/user/repository/mysql"
	_userPostgresRepository "github.com/jllanes-ss/avisos/account/user/repository/postgres"
)

func GetRepository(db *sql.DB) domain.UserRepository {
	sentry.CaptureMessage("obtube repo!")
	dv := reflect.ValueOf(db.Driver())
	switch dv.Type().String() {
	case "*pq.Driver":
		return _userPostgresRepository.NewPostgresUserRepository(db)
	default:
		return _userMysqlRepository.NewMysqlUserRepository(db)
	}

}
