package validators

import (
	"Assignment/constants"
	"Assignment/models"
	"Assignment/services/products"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func ValidateAddProductReq(ctx *gin.Context) (models.AddProductReq, error) {
	var req models.AddProductReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return models.AddProductReq{}, err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForAddProduct(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return models.AddProductReq{}, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	if req.Category != constants.ProductCategory.PREMIUM &&
		req.Category != constants.ProductCategory.REGULAR &&
		req.Category != constants.ProductCategory.BUDGET {
		return models.AddProductReq{}, errors.New("invalid category(must be premium, regular or budget)")
	}

	return req, nil
}

func ValidateListAllProductsReq(ctx *gin.Context) (string, error) {

	// validate category passed in query param
	category := ctx.Query("category")
	if category != "" {
		if category != constants.ProductCategory.PREMIUM &&
			category != constants.ProductCategory.REGULAR &&
			category != constants.ProductCategory.BUDGET {
			return "", errors.New("invalid category(must be premium, regular or budget)")
		}
	}

	return category, nil
}

func ValidateUpdateProductReq(ctx *gin.Context) (models.AddProductReq, string, error) {
	pid := ctx.Param("id")
	_, ok := products.ProductCatalog[pid]
	if !ok {
		return models.AddProductReq{}, pid, errors.New("no such product with id: " + pid)
	}

	var req models.AddProductReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return models.AddProductReq{}, pid, err
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: getRulesForUpdateProduct(),
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return models.AddProductReq{}, pid, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	// category validation
	if req.Category != "" {
		if req.Category != constants.ProductCategory.PREMIUM &&
			req.Category != constants.ProductCategory.REGULAR &&
			req.Category != constants.ProductCategory.BUDGET {
			return models.AddProductReq{}, pid, errors.New("invalid category(must be premium, regular or budget)")
		}
	}

	// atleast one field should be there to update
	if req.Category == "" && req.Name == "" && req.Price == 0.0 && req.Quantity == 0 {
		return models.AddProductReq{}, pid, errors.New("No field to update, add atleast one field")
	}

	return req, pid, nil
}

func ValidateDeleteProductReq(ctx *gin.Context) (string, error) {
	pid := ctx.Param("id")
	_, ok := products.ProductCatalog[pid]
	if !ok {
		return pid, errors.New("no such product with id: " + pid)
	}

	return pid, nil
}
