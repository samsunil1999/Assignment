package products

import (
	"Assignment/models"
	"fmt"

	"github.com/google/uuid"
)

type ProductImplementation struct{}

var ProductCatalog = make(map[string]models.ProductEntity)

func (p ProductImplementation) AddProduct(req models.AddProductReq) (res models.AddProductRes) {
	product := models.ProductEntity{
		Name:         req.Name,
		Category:     req.Category,
		Price:        req.Price,
		Availability: req.Quantity,
	}
	id := "pr_" + uuid.NewString()
	ProductCatalog[id] = product

	fmt.Println(id)
	fmt.Println(ProductCatalog)

	res.Status = "success"
	res.Message = "product added seccefully with id: " + id

	return res
}

func (p ProductImplementation) ListAllProduct(category string) models.ListAllProductsRes {

	var arrProd []models.FullProductDetails
	var recCount int
	for id, product := range ProductCatalog {
		if category != "" {
			if category == product.Category { // for listing particular category products
				arrProd = append(arrProd, models.FullProductDetails{
					ID:           id,
					Name:         product.Name,
					Category:     product.Category,
					Price:        product.Price,
					Availability: product.Availability,
				})
				recCount++
			}
		} else { // for listing all products
			arrProd = append(arrProd, models.FullProductDetails{
				ID:           id,
				Name:         product.Name,
				Category:     product.Category,
				Price:        product.Price,
				Availability: product.Availability,
			})
			recCount++
		}

	}

	resp := models.ListAllProductsRes{
		RecordCount: recCount,
		Data:        arrProd,
	}

	return resp

}

func (p ProductImplementation) UpdateProduct(req models.AddProductReq, pid string) (res models.FullProductDetails) {
	prodEntity := ProductCatalog[pid]

	// Note: can't be updated with zero values
	if req.Category != "" {
		prodEntity.Category = req.Category
	}
	if req.Name != "" {
		prodEntity.Name = req.Name
	}
	if req.Quantity != 0 {
		prodEntity.Availability = req.Quantity
	}
	if req.Price != 0.0 {
		prodEntity.Price = req.Price
	}

	// updating the ProductCatalog map
	ProductCatalog[pid] = prodEntity
	res = models.FullProductDetails{
		ID:           pid,
		Name:         prodEntity.Name,
		Category:     prodEntity.Category,
		Price:        prodEntity.Price,
		Availability: prodEntity.Availability,
	}

	return res
}

func (p ProductImplementation) DeleteProduct(pid string) (res models.AddProductRes) {
	delete(ProductCatalog, pid)

	res.Status = "success"
	res.Message = "product successfully deteled with id: " + pid
	return res
}
