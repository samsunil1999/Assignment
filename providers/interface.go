package providers

import (
	"Assignment/services/orders"
	"Assignment/services/products"
)

var (
	ProductSrv products.ProductInterface = products.ProductImplementation{}
	OrderSrv   orders.OrderInterface     = orders.OrderImplementation{}
)
