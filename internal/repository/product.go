package repository

import (
	"avito-shop-service/internal/db"
	"avito-shop-service/internal/models"
	"context"
)

// ProductRepository - интерфейс для репозитория продуктов
type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id int64) (*models.Product, error)
}

// productRepository - структура для работы с продуктами
type productRepository struct {
	db *db.PostgresDB
}

// NewProductRepository создает новый экземпляр репозитория
func NewProductRepository(db *db.PostgresDB) ProductRepository {
	return &productRepository{db: db}
}

// GetAllProducts реализует получение всех продуктов
func (r *productRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product

	// Пример: добавляем продукты вручную
	products = append(products, models.Product{ID: 1, Name: "Product 1", Price: 100})
	products = append(products, models.Product{ID: 2, Name: "Product 2", Price: 200})

	return products, nil
}

// GetProductByID реализует получение продукта по ID
func (r *productRepository) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	var product models.Product
	// Пример: возвращаем продукт без реального запроса
	return &product, nil
}
