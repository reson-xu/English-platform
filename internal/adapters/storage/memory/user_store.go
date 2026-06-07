package memory

import (
	"context"
	"sync"
	"time"

	"github.com/reson-xu/english-platform/internal/models/db"
	usermodule "github.com/reson-xu/english-platform/internal/modules/user"
)

type UserStore struct {
	mu      sync.RWMutex
	byID    map[string]db.User
	byEmail map[string]string
}

func NewUserStore() *UserStore {
	return &UserStore{
		byID:    map[string]db.User{},
		byEmail: map[string]string{},
	}
}

func (s *UserStore) Create(ctx context.Context, user db.User) (db.User, error) {
	if err := ctx.Err(); err != nil {
		return db.User{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.byEmail[user.Email]; exists {
		return db.User{}, usermodule.ErrEmailAlreadyExists
	}

	s.byID[user.ID] = user
	s.byEmail[user.Email] = user.ID

	return user, nil
}

func (s *UserStore) FindByEmail(ctx context.Context, email string) (db.User, error) {
	if err := ctx.Err(); err != nil {
		return db.User{}, err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	id, exists := s.byEmail[email]
	if !exists {
		return db.User{}, usermodule.ErrUserNotFound
	}

	return s.byID[id], nil
}

func (s *UserStore) UpdateLastLogin(ctx context.Context, id string, at time.Time) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.byID[id]
	if !exists {
		return usermodule.ErrUserNotFound
	}

	user.LastLoginAt = &at
	user.UpdatedAt = at
	s.byID[id] = user

	return nil
}
