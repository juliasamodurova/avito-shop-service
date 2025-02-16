-- Создание таблицы продуктов (мерча)
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,    -- Уникальный идентификатор продукта
    name VARCHAR(255) NOT NULL,           -- Название продукта
    price INT NOT NULL,                   -- Цена продукта в монетах
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата и время создания
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Дата и время последнего обновления
);

-- Вставка данных о товарах
INSERT INTO products (name, price) VALUES
('t-shirt', 80),
('cup', 20),
('book', 50),
('pen', 10),
('powerbank', 200),
('hoody', 300),
('umbrella', 200),
('socks', 10),
('wallet', 50),
('pink-hoody', 500);

-- Создание таблицы пользователей
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,               -- Уникальный идентификатор пользователя
    name VARCHAR(255) NOT NULL,                       -- Имя пользователя
    email VARCHAR(255) UNIQUE NOT NULL,               -- Электронная почта
    coins INT DEFAULT 1000,                           -- Количество монет у пользователя
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Дата и время регистрации
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Дата и время последнего обновления
);

-- Создание таблицы заказов
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,               -- Уникальный идентификатор заказа
    user_id INT,                                     -- Ссылка на пользователя
    total_amount INT NOT NULL,                       -- Общая сумма заказа в монетах
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Дата и время создания
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Дата и время последнего обновления
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE -- Связь с таблицей пользователей
);

-- Создание таблицы элементов заказа (связь продуктов и заказов)
CREATE TABLE order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,               -- Уникальный идентификатор элемента заказа
    order_id INT,                                    -- Ссылка на заказ
    product_id INT,                                  -- Ссылка на продукт
    quantity INT NOT NULL,                           -- Количество продуктов в заказе
    price INT NOT NULL,                              -- Цена одного продукта на момент заказа
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE, -- Связь с таблицей заказов
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE -- Связь с таблицей продуктов
);

-- Создание таблицы для перемещения монет между пользователями
CREATE TABLE coin_transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,               -- Уникальный идентификатор транзакции
    from_user_id INT,                                -- Ссылка на отправителя
    to_user_id INT,                                  -- Ссылка на получателя
    amount INT NOT NULL,                             -- Количество монет
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Дата и время транзакции
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,  -- Связь с отправителем
    FOREIGN KEY (to_user_id) REFERENCES users(id) ON DELETE CASCADE     -- Связь с получателем
);

-- Пример вставки тестового пользователя
INSERT INTO users (name, email, coins) VALUES
('John Doe', 'john.doe@example.com', 1000);

-- Пример заказа
INSERT INTO orders (user_id, total_amount) VALUES
(1, 100);

-- Пример элементов заказа
INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
(1, 1, 1, 80),  -- t-shirt
(1, 2, 1, 20);  -- cup

-- Пример транзакции монет
INSERT INTO coin_transactions (from_user_id, to_user_id, amount) VALUES
(1, 2, 100);  -- Иван Иванов отправил 100 монет другому пользователю
