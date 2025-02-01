package product

import (
	"ecobuy/entities"
	"ecobuy/repositories/product"
)

type ProductServiceInterface interface {
	GetProducts(category string, page int, limit int) ([]entities.Product, error)
	GetProductByID(id int) (*entities.Product, error)
	GetImpactByProductID(productID uint) (*entities.ImpactData, error)
}

type ProductService struct {
	ProductRepo product.ProductRepositoryInterface
}

func NewProductService(pr product.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		ProductRepo: pr,
	}
}

func (ps *ProductService) GetProducts(category string, page int, limit int) ([]entities.Product, error) {
	return ps.ProductRepo.GetProducts(category, page, limit)
}

func (ps *ProductService) GetProductByID(id int) (*entities.Product, error) {
	return ps.ProductRepo.GetProductByID(id)
}

func (s *ProductService) GetImpactByProductID(productID uint) (*entities.ImpactData, error) {
	return s.ProductRepo.GetImpactByProductID(productID)
}
