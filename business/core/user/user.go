// Package user provides an example of a core business API. Right now these
// calls are just wrapping the data/data layer. But at some point you will
// want auditing or something that isn't specific to the data/store layer.
package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jcsix694/service3-video/business/data/store/user"
	"github.com/jcsix694/service3-video/business/sys/auth"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Core manages the set of APIs for user access.
type Core struct {
	log  *zap.SugaredLogger
	user user.Store
}

// NewCore constructs a core for user api access.
func NewCore(log *zap.SugaredLogger, db *sqlx.DB) Core {
	return Core{
		log:  log,
		user: user.NewStore(log, db),
	}
}

// Create inserts a new user into the database.
func (c Core) Create(ctx context.Context, nu user.NewUser, now time.Time) (user.User, error) {
	usr, err := c.user.Create(ctx, nu, now)

	if err != nil {
		return usr, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}

// Update updates a user from the database.
func (c Core) Update(ctx context.Context, claims auth.Claims, userID string, uu user.UpdateUser, now time.Time) error {

	if err := c.user.Update(ctx, claims, userID, uu, now); err != nil {
		fmt.Errorf("update: %w", err)
	}

	return nil
}

// Delete deletes a user from the database.
func (c Core) Delete(ctx context.Context, claims auth.Claims, userID string) error {

	if err := c.user.Delete(ctx, claims, userID); err != nil {
		fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing users from the database.
func (c Core) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]user.User, error) {

	users, err := c.user.Query(ctx, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	return users, nil
}

// Query retrieves a user by id.
func (c Core) QueryById(ctx context.Context, claims auth.Claims, userID string) (user.User, error) {

	user, err := c.user.QueryByID(ctx, claims, userID)

	if err != nil {
		return user, fmt.Errorf("query: %w", err)
	}
	return user, nil
}

// Query retrieves a user by email.
func (c Core) QueryByEmail(ctx context.Context, claims auth.Claims, email string) (user.User, error) {

	user, err := c.user.QueryByEmail(ctx, claims, email)

	if err != nil {
		return user, fmt.Errorf("query: %w", err)
	}
	return user, nil
}

// Authenticate finds a user by their email and verifies their password. On Il success it returns a Claims User representing this user. The claims can be
// used to generate a token for future authentication.
func (c Core) Authenticate(ctx context.Context, now time.Time, email, password string) (auth.Claims, error) {
	claims, err := c.user.Authenticate(ctx, now, email, password)
	if err != nil {
		return auth.Claims{}, fmt.Errorf("query: %w", err)
	}
	// PERFORM POST BUSINESS OPERATIONS,
	return claims, nil
}
