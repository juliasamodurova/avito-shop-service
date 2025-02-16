package db

import (
	"database/sql"
	"fmt"


	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// Структура для подключения к PostgreSQL
type PostgresDB struct {
	Conn *sql.DB
}

// Функция для создания подключения к PostgreSQL
func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=youruser dbname=yourdbname password=yourpassword sslmode=disable" // Замените на свои данные
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}

	return &PostgresDB{Conn: db}, nil
}

// Закрытие подключения
func (db *PostgresDB) Close() error {
	return db.Conn.Close()
}
