package products

import (
	"Assignment/controllers"
	"Assignment/providers"
	"Assignment/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProductHandler(ctx *gin.Context) {
	req, err := validators.ValidateAddProductReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res := providers.ProductSrv.AddProduct(req)
	controllers.ReturnJsonStruct(ctx, res)
}

func ListAllProductsHandler(ctx *gin.Context) {
	category, err := validators.ValidateListAllProductsReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	res := providers.ProductSrv.ListAllProduct(category)
	controllers.ReturnJsonStruct(ctx, res)
}

func UpdateProductHandler(ctx *gin.Context) {
	req, pid, err := validators.ValidateUpdateProductReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res := providers.ProductSrv.UpdateProduct(req, pid)
	controllers.ReturnJsonStruct(ctx, res)
}

func DeleteProductHandler(ctx *gin.Context) {
	pid, err := validators.ValidateDeleteProductReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res := providers.ProductSrv.DeleteProduct(pid)
	controllers.ReturnJsonStruct(ctx, res)
}
