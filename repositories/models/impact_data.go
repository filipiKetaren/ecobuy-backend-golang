package models

import (
	"ecobuy/entities"
	"time"
)

type ImpactData struct {
	ID                    uint   `gorm:"primaryKey"`
	ProductID             uint   `gorm:"not null"`
	CarbonFootprint       string `gorm:"type:varchar(255)"`
	WaterSaving           string `gorm:"type:varchar(255)"`
	PlasticWasteReduction string `gorm:"type:varchar(255)"`
	CreatedAt             time.Time
	UpdatedAt             time.Time

	Product Product `gorm:"foreignKey:ProductID"`
}

func FromEntitiesImpactData(impact entities.ImpactData) ImpactData {
	return ImpactData{
		ID:                    impact.ID,
		ProductID:             impact.ProductID,
		CarbonFootprint:       impact.CarbonFootprint,
		WaterSaving:           impact.CarbonFootprint,
		PlasticWasteReduction: impact.PlasticWasteReduction,
		CreatedAt:             impact.CreatedAt,
		UpdatedAt:             impact.UpdatedAt,
	}
}

func (impact ImpactData) ToEntities() entities.ImpactData {
	return entities.ImpactData{
		ID:                    impact.ID,
		ProductID:             impact.ProductID,
		CarbonFootprint:       impact.CarbonFootprint,
		WaterSaving:           impact.CarbonFootprint,
		PlasticWasteReduction: impact.PlasticWasteReduction,
		CreatedAt:             impact.CreatedAt,
		UpdatedAt:             impact.UpdatedAt,
	}
}
