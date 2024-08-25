package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryInterface interface {
	GetByEmail(email string) (*models.Auth, error)
}

type AuthRepository struct {
	*sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) GetByEmail(email string) (*models.Auth, error) {
	result := models.Auth{}
	query := `SELECT id, email, password, role from public.users WHERE email = $1`
	err := r.Get(&result, query, email)
	if err != nil {
		return nil, err
	}

	return &result, nil
}