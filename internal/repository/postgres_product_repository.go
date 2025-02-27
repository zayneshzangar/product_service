package repository

import (
	"context"
	"database/sql"
	"product_service/internal/entity"
)

// PostgresProductRepository реализует интерфейс ProductRepository для Postgres
type PostgresProductRepository struct {
	db *sql.DB
}

// NewPostgresProductRepository создаёт новый Postgres-репозиторий
func NewPostgresProductRepository(db *sql.DB) ProductRepository {
	return &PostgresProductRepository{db: db}
}

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

// STOCK methods
func (r *PostgresProductRepository) BeginTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *PostgresProductRepository) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	product := &entity.Product{}
	err := r.db.QueryRowContext(ctx, "SELECT id, stock FROM products WHERE id = $1", productID).
		Scan(&product.ID, &product.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *PostgresProductRepository) ReserveStock(ctx context.Context, tx *sql.Tx, orderID, productID int64, quantity int) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO reserved_stock (order_id, product_id, quantity) VALUES ($1, $2, $3)", orderID, productID, quantity)
	return err
}

func (r *PostgresProductRepository) ClearExpiredReservations(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM reserved_stock WHERE created_at < NOW() - INTERVAL '15 minutes'`)
	return err
}
