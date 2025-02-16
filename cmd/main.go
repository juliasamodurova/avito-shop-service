package repository

import (
	"context"
	"avito-shop-service/internal/app"
	"avito-shop-service/internal/db"
	"fmt"
)

// Структура для репозитория продуктов с использованием PostgreSQL
type PostgresUserRepositoryProductService struct {
	DB *db.PostgresDB
}

// Новый репозиторий для работы с продуктами
func NewPostgresUserRepositoryProductService(DB *db.PostgresDB) *PostgresUserRepositoryProductService {
	return &PostgresUserRepositoryProductService{DB: DB}
}

// Получение всех продуктов из базы данных
func (s *PostgresUserRepositoryProductService) GetAllProducts(ctx context.Context) ([]app.Product, error) {
	query := "SELECT id, name, description, price FROM products"
	rows, err := s.DB.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch products: %v", err)
	}
	defer rows.Close()

	var products []app.Product
	for rows.Next() {
		var p app.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price); err != nil {
			return nil, fmt.Errorf("unable to scan product: %v", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return products, nil
}
