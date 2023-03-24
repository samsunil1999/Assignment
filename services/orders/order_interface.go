package orders

import "Assignment/models"

type OrderInterface interface {
	CreateOrder(req models.CreateOrderReq) (models.CreateOrderRes, error)
	ListAllOrders() (res models.ListAllOrdersRes)
	GetOrdersById(pid string) models.CreateOrderRes
	UpdateOrderById(pid string) (models.UpdateOrderStatus, error)
	CancelOrderById(pid string) (models.UpdateOrderStatus, error)
	// UpdateProduct(req models.AddProductReq, pid string) (res models.FullProductDetails)
	// DeleteProduct(pid string) (res models.AddProductRes)
}
