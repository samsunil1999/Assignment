package models

// ------------------------------------------
// Entity
type ProductEntity struct {
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
}

// ------------------------------------------
// Add product
type AddProductReq struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type AddProductRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ------------------------------------------
// List all product
type FullProductDetails struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
}
type ListAllProductsRes struct {
	RecordCount int                  `json:"record_count"`
	Data        []FullProductDetails `json:"products"`
}

// ------------------------------------------
