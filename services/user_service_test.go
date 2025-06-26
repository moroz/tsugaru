package services_test

import (
	"context"
	"database/sql"
	"oauth-provider/config"
	"oauth-provider/db/queries"
	"oauth-provider/services"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var TEST_DATABASE_URL = config.MustGetenv("TEST_DATABASE_URL")

func initDB(ctx context.Context) (queries.DBTX, error) {
	return pgx.Connect(ctx, TEST_DATABASE_URL)
}

func TestAuthenticateUserByEmailPassword(t *testing.T) {
	db, err := initDB(t.Context())
	require.NoError(t, err)
	srv := services.NewUserService(db)

	t.Run("authenticates user with valid email and password", func(t *testing.T) {
		emails := []string{"user@example.com", "User@Example.Com", "USER@EXAMPLE.COM"}

		for _, email := range emails {
			actual, err := srv.AuthenticateUserByEmailPassword(t.Context(), email, "hunter2")
			assert.NoError(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, actual.Email, "user@example.com")
		}
	})

	t.Run("returns ErrPasswordDisabled when user has no password", func(t *testing.T) {
		actual, err := srv.AuthenticateUserByEmailPassword(t.Context(), "nopassword@example.com", "hunter2")
		assert.ErrorIs(t, err, services.ErrPasswordDisabled)
		assert.Nil(t, actual)
	})

	t.Run("returns ErrInvalidPassword with invalid password", func(t *testing.T) {
		actual, err := srv.AuthenticateUserByEmailPassword(t.Context(), "user@example.com", "invalid")
		assert.ErrorIs(t, err, services.ErrInvalidPassword)
		assert.Nil(t, actual)
	})

	t.Run("returns error with non-existent user", func(t *testing.T) {
		actual, err := srv.AuthenticateUserByEmailPassword(t.Context(), "non-existent", "hunter2")
		assert.ErrorIs(t, err, sql.ErrNoRows)
		assert.Nil(t, actual)
	})
}
