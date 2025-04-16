package repository

import (
	"database/sql"
	"fmt"
	"product_service/internal/entity"
)

// PostgresProductRepository реализует интерфейс ProductRepository для Postgres
type PostgresProductRepository struct {
	db *sql.DB
}

// // NewPostgresProductRepository создаёт новый Postgres-репозиторий
// func NewPostgresProductRepository(db *sql.DB) ProductRepository {
// 	return &PostgresProductRepository{db: db}
// }

// Методы работы с продуктами

func (r *PostgresProductRepository) Create(product *entity.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4)",
		product.Name, product.Description, product.Price, product.Stock)
	return err
}

func (r *PostgresProductRepository) GetByID(id int64) (*entity.Product, error) {
	var product entity.Product
	err := r.db.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id=$1", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	return &product, err
}

func (r *PostgresProductRepository) GetAll() ([]*entity.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *PostgresProductRepository) Update(product *entity.Product) error {
	_, err := r.db.Exec("UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$5",
		product.Name, product.Description, product.Price, product.Stock, product.ID)
	return err
}

func (r *PostgresProductRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

func (r *PostgresProductRepository) Close() error {
	return r.db.Close()
}

func (r *PostgresProductRepository) UpdateStock(productId int64, quantity int64) error {
	result, err := r.db.Exec("UPDATE products SET stock = stock - $1, updated_at = NOW() WHERE id = $2 AND stock >= $1", quantity, productId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("insufficient stock for product %d", productId)
	}

	return nil
}
