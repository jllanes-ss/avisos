package domain

import (
	"github.com/google/uuid"
)

const (
	Enabled  = "ENABLED"  // the user is enabled.
	Disabled = "DISABLED" // the user is disabled
	Pending  = "PENDING"  // The user is pending when the user use the registration usercase that need tha a user administrator approve the registration
)

type User struct {
	ID       uuid.UUID `db:"uid" json:"uid"`
	Name     string    `db:"name" json:"name"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"`
}
