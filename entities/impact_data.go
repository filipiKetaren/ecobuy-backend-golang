package entities

import "time"

type ImpactData struct {
	ID                    uint
	ProductID             uint
	CarbonFootprint       string
	WaterSaving           string
	PlasticWasteReduction string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
