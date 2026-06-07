package user

import (
	"context"
	"time"

	"github.com/reson-xu/english-platform/internal/models/db"
)

type Repository interface {
	Create(ctx context.Context, user db.User) (db.User, error)
	FindByEmail(ctx context.Context, email string) (db.User, error)
	UpdateLastLogin(ctx context.Context, id string, at time.Time) error
}
