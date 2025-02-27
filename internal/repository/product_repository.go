package repository

import (
	"context"
	"database/sql"
	"product_service/internal/entity"
)

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
	Create(product *entity.Product) error
	GetByID(id int64) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id int64) error
	BeginTransaction() (*sql.Tx, error)
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
	ReserveStock(ctx context.Context, tx *sql.Tx, orderID, productID int64, quantity int) error
	ClearExpiredReservations(ctx context.Context) error
}
