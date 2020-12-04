package repositories

import (
	"database/sql"
	"golang_repository_restapi/models"
)

// UserRepository implements models.UserRepository
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository ..
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// All ..
func (r *UserRepository) All() (*[]models.User, error) {
	var users []models.User
	return &users, nil
}
