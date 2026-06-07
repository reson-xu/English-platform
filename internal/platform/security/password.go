package security

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	passwordHashAlgorithm  = "pbkdf2_sha256"
	passwordHashIterations = 120_000
	passwordHashSaltBytes  = 16
	passwordHashKeyBytes   = 32
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (s *PasswordService) HashPassword(password string) (string, error) {
	salt := make([]byte, passwordHashSaltBytes)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("read salt: %w", err)
	}

	key := pbkdf2SHA256(
		[]byte(password),
		salt,
		passwordHashIterations,
		passwordHashKeyBytes,
	)

	return strings.Join(
		[]string{
			passwordHashAlgorithm,
			strconv.Itoa(passwordHashIterations),
			base64.RawURLEncoding.EncodeToString(salt),
			base64.RawURLEncoding.EncodeToString(key),
		},
		"$",
	), nil
}

func (s *PasswordService) ComparePassword(password string, encoded string) (bool, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 || parts[0] != passwordHashAlgorithm {
		return false, errors.New("unsupported password hash")
	}

	iterations, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("parse password hash iterations: %w", err)
	}

	salt, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return false, fmt.Errorf("decode password hash salt: %w", err)
	}

	expected, err := base64.RawURLEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("decode password hash key: %w", err)
	}

	actual := pbkdf2SHA256([]byte(password), salt, iterations, len(expected))
	return subtle.ConstantTimeCompare(actual, expected) == 1, nil
}

func pbkdf2SHA256(password []byte, salt []byte, iterations int, keyLen int) []byte {
	derived := []byte{}
	var blockIndex uint32 = 1

	for len(derived) < keyLen {
		block := pbkdf2Block(password, salt, iterations, blockIndex)
		derived = append(derived, block...)
		blockIndex++
	}

	return derived[:keyLen]
}

func pbkdf2Block(password []byte, salt []byte, iterations int, blockIndex uint32) []byte {
	mac := hmac.New(sha256.New, password)
	mac.Write(salt)
	mac.Write([]byte{
		byte(blockIndex >> 24),
		byte(blockIndex >> 16),
		byte(blockIndex >> 8),
		byte(blockIndex),
	})

	u := mac.Sum(nil)
	result := make([]byte, len(u))
	copy(result, u)

	for range iterations - 1 {
		mac = hmac.New(sha256.New, password)
		mac.Write(u)
		u = mac.Sum(nil)

		for i := range result {
			result[i] ^= u[i]
		}
	}

	return result
}
