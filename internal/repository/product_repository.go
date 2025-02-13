package repository

import "product_service/internal/entity"

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
    Create(product *entity.Product) error
    GetByID(id int64) (*entity.Product, error)
    GetAll() ([]*entity.Product, error)
    Update(product *entity.Product) error
    Delete(id int64) error
}
