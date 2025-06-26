package services

import (
	"context"
	"errors"
	"fmt"
	"oauth-provider/db/queries"

	"github.com/alexedwards/argon2id"
)

var ErrPasswordDisabled = errors.New("password authentication disabled")
var ErrInvalidPassword = errors.New("invalid password")

type UserService struct {
	db queries.DBTX
}

func NewUserService(db queries.DBTX) *UserService {
	return &UserService{db}
}

func (s *UserService) AuthenticateUserByEmailPassword(ctx context.Context, email, password string) (*queries.User, error) {
	user, err := queries.New(s.db).GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash == nil {
		return nil, fmt.Errorf("%w for user %s", ErrPasswordDisabled, user.Email)
	}

	match, _, err := argon2id.CheckHash(password, *user.PasswordHash)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
