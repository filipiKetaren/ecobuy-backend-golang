package product

import (
	"ecobuy/entities"
	"ecobuy/repositories/models"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetProducts(category string, page int, limit int) ([]entities.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProducts(category string, page int, limit int) ([]entities.Product, error) {
	var products []models.Product
	query := pr.db

	// Jika kategori diberikan, tambahkan filter
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Pagination menggunakan Offset dan Limit
	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}

	var entityProducts []entities.Product
	for _, product := range products {
		entityProducts = append(entityProducts, product.ToEntities())
	}

	return entityProducts, nil
}
