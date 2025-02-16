package models

// Структура для продукта
type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price int `json:"price"`
}

// Структура для пользователя
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
