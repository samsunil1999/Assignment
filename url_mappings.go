package main

import (
	"Assignment/controllers/orders"
	"Assignment/controllers/products"
)

func mapUrls() {
	// product endpoints
	product := router.Group("/product")
	product.POST("/add", products.AddProductHandler)
	product.GET("/list-all", products.ListAllProductsHandler)
	product.PUT("/update/:id", products.UpdateProductHandler)
	product.DELETE("/delete/:id", products.DeleteProductHandler)

	// order endpoints
	order := router.Group("/order")
	order.POST("/create", orders.CreateOrderHandler)
	order.GET("/list-all", orders.GetAllOrdersHandler)
	order.GET("/:id", orders.GetOrderByIdHandler)
	order.PUT("/update-status/:id", orders.UpdateOrderStatusHandler)
	order.POST("/cancel/:id", orders.CancelOrderHandler)
}
