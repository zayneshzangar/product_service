package usecase

import (
	"errors"
	"time"

	"product_service/internal/entity"
	"product_service/internal/repository"
)

// ProductUseCase — интерфейс бизнес-логики
type ProductUseCase interface {
	CreateProduct(name, description string, price float64, stock int) (*entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id int64) error
}

// productUseCase — реализация use case
type productUseCase struct {
	productRepo repository.ProductRepository
}

// NewProductUseCase создает новый use case
func NewProductUseCase(productRepo repository.ProductRepository) ProductUseCase {
	return &productUseCase{productRepo: productRepo} // ВОЗВРАЩАЕМ УКАЗАТЕЛЬ
}

// CreateProduct создает новый продукт
func (u *productUseCase) CreateProduct(name, description string, price float64, stock int) (*entity.Product, error) {
	if name == "" || price <= 0 || stock < 0 {
		return nil, errors.New("invalid product data")
	}

	product := &entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := u.productRepo.Create(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductByID получает продукт по ID
func (u *productUseCase) GetProductByID(id int64) (*entity.Product, error) {
	return u.productRepo.GetByID(id)
}

// GetAllProducts получает все продукты
func (u *productUseCase) GetAllProducts() ([]*entity.Product, error) {
	return u.productRepo.GetAll()
}

// UpdateProduct обновляет продукт
func (u *productUseCase) UpdateProduct(product *entity.Product) error {
	if product.ID == 0 {
		return errors.New("invalid product ID")
	}
	product.UpdatedAt = time.Now()
	return u.productRepo.Update(product)
}

// DeleteProduct удаляет продукт
func (u *productUseCase) DeleteProduct(id int64) error {
	return u.productRepo.Delete(id)
}
