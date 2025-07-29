package infrastructure

import (
	"task_manager/domain"

	"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

// NewPasswordService creates a new password service instance
func NewPasswordService() domain.IPasswordService {
	return &passwordService{}
}

func (p *passwordService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (p *passwordService) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
