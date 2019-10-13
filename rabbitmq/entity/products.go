package entity

//Product entity 
type Product struct {
	ID       int
	Thumb    []byte
	Name     string
	Group    string
	Quantity int
	Price    float64
}