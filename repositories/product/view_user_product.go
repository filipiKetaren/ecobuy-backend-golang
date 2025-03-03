package product

import (
	"ecobuy/entities"
	"ecobuy/repositories/models"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetProducts(category string, page int, limit int) ([]entities.Product, error)
	GetProductByID(id int) (*entities.Product, error)
	GetImpactByProductID(productID uint) (*entities.ImpactData, error)
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

func (pr *ProductRepository) GetProductByID(id int) (*entities.Product, error) {
	var product models.Product
	err := pr.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	entityProduct := product.ToEntities()
	return &entityProduct, nil
}

func (r *ProductRepository) GetImpactByProductID(productID uint) (*entities.ImpactData, error) {
	var impact models.ImpactData
	err := r.db.Where("product_id = ?", productID).First(&impact).Error
	if err != nil {
		return nil, err
	}
	entityImpact := impact.ToEntities()
	return &entityImpact, nil
}
