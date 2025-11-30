package security

import (
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
	"golang.org/x/crypto/bcrypt"
)

type BCryptPasswordHasher struct{}

func NewBCryptPasswordHasher() *BCryptPasswordHasher {
	logger.Log.Info("Password Hasher 'BCryptPasswordHasher' setted...")
	return &BCryptPasswordHasher{}
}

func (b *BCryptPasswordHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashValue := string(hashed)
	logger.Log.Debug("Hash generated", "token", hashValue)
	return hashValue, nil
}

func (b *BCryptPasswordHasher) Compare(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
