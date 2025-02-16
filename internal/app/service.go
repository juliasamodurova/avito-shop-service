package app

import (
	"avito-shop-service/internal/models" // Добавить импорт для моделей
	"avito-shop-service/internal/repository"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Секретный ключ для подписи токена
var secretKey = []byte("your_secret_key")

// Структура для токена
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Структура сервиса для продуктов
type ProductService struct {
	repo repository.ProductRepository
}

// Структура сервиса для пользователей
type UserService struct {
	repo repository.UserRepository
}

// Преобразует models.Product в app.Product с приведением типов
func ConvertToAppProduct(p models.Product) Product {
	return Product{
		ID:    int(p.ID), // Преобразуем int64 в int
		Name:  p.Name,
		Price: p.Price,
	}
}

// Создание JWT
func CreateJWT(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Токен истекает через 24 часа
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Хэширование пароля с использованием bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Проверка пароля
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Регистрация нового пользователя
func (s *UserService) RegisterUser(username, password string) (string, error) {
	// Проверка существующего пользователя
	user, err := s.repo.GetUserByUsername(username)
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return "", err
	}

	if user != nil {
		return "", errors.New("пользователь с таким именем уже существует")
	}

	// Хэширование пароля
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return "", err
	}

	// Создание нового пользователя
	newUser := &models.User{
		Username: username,
		Password: hashedPassword,
	}

	// Сохранение пользователя в базе данных
	err = s.repo.SaveUser(newUser)
	if err != nil {
		return "", err
	}

	return newUser.ID, nil
}

// Авторизация пользователя
func (s *UserService) AuthenticateUser(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// Проверка пароля
	if !CheckPasswordHash(password, user.Password) {
		return "", errors.New("неправильный пароль")
	}

	// Создание JWT
	token, err := CreateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Новый сервис для продуктов
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Новый сервис для пользователей
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Получение продукта по ID
func (s *ProductService) GetProductByID(ctx context.Context, id int64) (*Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	appProduct := ConvertToAppProduct(*product)
	return &appProduct, nil
}

// Получение всех продуктов
func (s *ProductService) GetAllProducts(ctx context.Context) ([]Product, error) {
	products, err := s.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	var appProducts []Product
	for _, p := range products {
		appProducts = append(appProducts, ConvertToAppProduct(p))
	}

	return appProducts, nil
}
