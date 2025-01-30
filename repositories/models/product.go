package models

import (
	"ecobuy/entities"
	"time"
)

type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Category    string    `gorm:"type:varchar(255);not null"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Stock       int       `gorm:"type:int;not null"`
	ImpactData  string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func FromEntitiesProduct(product entities.Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
		Stock:       product.Stock,
		ImpactData:  product.ImpactData,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (p Product) ToEntities() entities.Product {
	return entities.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		ImpactData:  p.ImpactData,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
