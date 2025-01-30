package entities

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Category    string
	Price       float64
	Stock       int
	ImpactData  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
