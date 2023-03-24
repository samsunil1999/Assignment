package models

// ------------------------------------------
// Entity
type OrderedProductData struct {
	ProductID       string `json:"product_id"`
	ProductQuantity int    `json:"product_quantity"`
}

type OrderEntity struct {
	Status        string               `json:"status"`
	Product       []OrderedProductData `json:"product_data"`
	OrderAmount   float64              `json:"order_amount"`
	Discount      bool                 `json:"discount_eligibility"`
	PayableAmount float64              `json:"payable_amount"`
	DispatchDate  string               `json:"dispatch_date,omitempty"`
}

// ------------------------------------------
// Create Order
type CreateOrderReq struct {
	Inventory []OrderedProductData `json:"inventory"`
}

type CreateOrderRes struct {
	ID            string               `json:"id"`
	Inventory     []OrderedProductData `json:"inventory"`
	OrderAmount   float64              `json:"order_amount"`
	Discount      bool                 `json:"discount"`
	PayableAmount float64              `json:"payable_amount"`
	Status        string               `json:"status"`
	DispatchDate  string               `json:"dispatch_date,omitempty"`
}

// ------------------------------------------
// List Orders
type OrderListDetails struct {
	ID            string  `json:"id"`
	Status        string  `json:"status"`
	PayableAmount float64 `json:"payable_amount"`
	DispatchDate  string  `json:"dispatch_date,omitempty"`
}

type ListAllOrdersRes struct {
	RecordCount int                `json:"record_count"`
	Data        []OrderListDetails `json:"orders"`
}

// ------------------------------------------
// Update Order status
type UpdateOrderStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
