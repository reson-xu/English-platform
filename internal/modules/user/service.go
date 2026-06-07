package user

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/reson-xu/english-platform/internal/models/db"
	userdto "github.com/reson-xu/english-platform/internal/models/dto/user"
	"github.com/reson-xu/english-platform/internal/platform/constants"
	apperrors "github.com/reson-xu/english-platform/internal/platform/errors"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(password string, encoded string) (bool, error)
}

type AccessTokenIssuer interface {
	GenerateAccessToken(userID string, email string, role string) (string, error)
}

type Service struct {
	users     Repository
	passwords PasswordHasher
	tokens    AccessTokenIssuer
}

func NewService(users Repository, passwords PasswordHasher, tokens AccessTokenIssuer) *Service {
	return &Service{
		users:     users,
		passwords: passwords,
		tokens:    tokens,
	}
}

func (s *Service) Register(ctx context.Context, req userdto.RegisterReq) (userdto.RegisterResp, error) {
	email, err := normalizeEmail(req.Email)
	if err != nil {
		return userdto.RegisterResp{}, err
	}

	nickname, err := normalizeNickname(req.Nickname)
	if err != nil {
		return userdto.RegisterResp{}, err
	}

	if err := validatePassword(req.Password); err != nil {
		return userdto.RegisterResp{}, err
	}

	if _, err := s.users.FindByEmail(ctx, email); err == nil {
		return userdto.RegisterResp{}, ErrEmailAlreadyExists
	} else if !errors.Is(err, ErrUserNotFound) {
		return userdto.RegisterResp{}, err
	}

	passwordHash, err := s.passwords.HashPassword(req.Password)
	if err != nil {
		return userdto.RegisterResp{}, apperrors.Wrap(err, apperrors.CodeInternal, "failed to hash password")
	}

	now := time.Now().UTC()
	user := db.User{
		ID:           newID(),
		Email:        email,
		PasswordHash: passwordHash,
		Nickname:     nickname,
		Role:         constants.UserRoleStudent,
		Status:       constants.UserStatusActive,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	created, err := s.users.Create(ctx, user)
	if err != nil {
		return userdto.RegisterResp{}, err
	}

	token, err := s.tokens.GenerateAccessToken(created.ID, created.Email, created.Role)
	if err != nil {
		return userdto.RegisterResp{}, apperrors.Wrap(err, apperrors.CodeInternal, "failed to generate access token")
	}

	return authResp(created, token), nil
}

func (s *Service) Login(ctx context.Context, req userdto.LoginReq) (userdto.LoginResp, error) {
	email, err := normalizeEmail(req.Email)
	if err != nil {
		return userdto.LoginResp{}, err
	}

	if strings.TrimSpace(req.Password) == "" {
		return userdto.LoginResp{}, invalidArgument("password is required")
	}

	user, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return userdto.LoginResp{}, ErrInvalidCredentials
		}

		return userdto.LoginResp{}, err
	}

	if user.Status != constants.UserStatusActive {
		return userdto.LoginResp{}, ErrUserDisabled
	}

	ok, err := s.passwords.ComparePassword(req.Password, user.PasswordHash)
	if err != nil {
		return userdto.LoginResp{}, apperrors.Wrap(err, apperrors.CodeInternal, "failed to compare password")
	}
	if !ok {
		return userdto.LoginResp{}, ErrInvalidCredentials
	}

	now := time.Now().UTC()
	if err := s.users.UpdateLastLogin(ctx, user.ID, now); err != nil {
		return userdto.LoginResp{}, fmt.Errorf("update last login: %w", err)
	}
	user.LastLoginAt = &now

	token, err := s.tokens.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return userdto.LoginResp{}, apperrors.Wrap(err, apperrors.CodeInternal, "failed to generate access token")
	}

	return authResp(user, token), nil
}

func authResp(user db.User, token string) userdto.RegisterResp {
	return userdto.RegisterResp{
		User: userdto.AuthUserResp{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
		},
		AccessToken: token,
	}
}

func normalizeEmail(value string) (string, error) {
	email := strings.ToLower(strings.TrimSpace(value))
	if email == "" {
		return "", invalidArgument("email is required")
	}

	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return "", invalidArgument("email is invalid")
	}

	return email, nil
}

func normalizeNickname(value string) (string, error) {
	nickname := strings.TrimSpace(value)
	if nickname == "" {
		return "", invalidArgument("nickname is required")
	}
	if len(nickname) < 2 || len(nickname) > 40 {
		return "", invalidArgument("nickname must be between 2 and 40 characters")
	}

	return nickname, nil
}

func validatePassword(value string) error {
	if value == "" {
		return invalidArgument("password is required")
	}
	if len(value) < 8 || len(value) > 128 {
		return invalidArgument("password must be between 8 and 128 characters")
	}

	return nil
}

func newID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		now := time.Now().UTC().UnixNano()
		return hex.EncodeToString([]byte(strconv.FormatInt(now, 10)))
	}

	return hex.EncodeToString(bytes)
}
