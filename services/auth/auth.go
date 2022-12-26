package auth

import (
	"crypto"
	"crypto/hmac"
	"fmt"
	"server-poc/pkg/models"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	DB            *gorm.DB
	Secret        []byte        // Secret used for generating and authenticating tokens
	TokenLifetime time.Duration // Lifetime of an authenticaton token
}

// Login verifies the credentials and generates a JWT token
func (s *Service) Login(username, password string) (string, error) {
	// Verify credentials
	var u models.User
	if err := u.FindByCreds(s.DB, username, s.hash(password)); err != nil {
		return "", fmt.Errorf("credentials don't match any accounts in the database: %w", err)
	}

	// Generate a token
	tok, err := createToken(u, s.TokenLifetime, s.Secret)
	if err != nil {
		return "", fmt.Errorf("failed to generate a jwt token: %w", err)
	}

	return tok, nil
}

// Register adds user to the databse
func (s *Service) Register(username, password string) error {
	// Check for collisions
	var u models.User
	if err := u.FindByUsername(s.DB, username); err == nil {
		return fmt.Errorf(`user named "%s" already exists`, username)
	}

	// Save in the database
	u = models.User{Username: username, PasswordHash: s.hash(password)}
	if err := u.Save(s.DB); err != nil {
		return fmt.Errorf(`failed to save user in the database: %w`, err)
	}

	return nil
}

// Authenticate verifies the token and returns embedded user data
func (s *Service) Authenticate(tok string) (models.User, error) {
	u, err := verifyToken(tok, s.Secret)
	if err != nil {
		return models.User{}, fmt.Errorf("token invalid: %w", err)
	}
	return u, nil
}

// hash generates a sha256 hmac hash of the input string
func (s *Service) hash(data string) string {
	mac := hmac.New(crypto.SHA256.New, s.Secret)
	mac.Write([]byte(data))
	return string(mac.Sum(nil))
}
