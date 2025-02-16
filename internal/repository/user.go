package repository

import (
	"avito-shop-service/internal/db"
	"avito-shop-service/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// Ошибка, если пользователь не найден
var ErrUserNotFound = errors.New("пользователь не найден")

// Интерфейс для работы с пользователями
type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	SaveUser(user *models.User) error
}

// Реализация репозитория для работы с пользователями через PostgreSQL
type PostgresUserRepository struct {
	DB *db.PostgresDB
}

// Метод для получения пользователя по имени пользователя
func (r *PostgresUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	// Используем QueryRowContext вместо QueryRow
	err := r.DB.Conn.QueryRowContext(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return &user, nil
}

// Метод для сохранения нового пользователя в базе данных
func (r *PostgresUserRepository) SaveUser(user *models.User) error {
	// Используем ExecContext вместо Exec
	_, err := r.DB.Conn.ExecContext(context.Background(), "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", user.ID, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("error saving user: %v", err)
	}
	return nil
}
