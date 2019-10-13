package entity

import "time"

//Product entity
type Product struct {
	ID        int    `json:id`
	Thumb     []byte `json:thumb`
	Name      string
	Group     string
	Quantity  int
	Price     float64
	CreatedAt time.Time
}
