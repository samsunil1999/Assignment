package products

import "Assignment/models"

type ProductInterface interface {
	AddProduct(req models.AddProductReq) (res models.AddProductRes)
	ListAllProduct(category string) (res models.ListAllProductsRes)
	UpdateProduct(req models.AddProductReq, pid string) (res models.FullProductDetails)
	DeleteProduct(pid string) (res models.AddProductRes)
}
