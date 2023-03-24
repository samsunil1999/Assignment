package orders

import (
	"Assignment/controllers"
	"Assignment/providers"
	"Assignment/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(ctx *gin.Context) {
	req, err := validators.ValidateCreateOrderReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res, err := providers.OrderSrv.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	controllers.ReturnJsonStruct(ctx, res)
}

func GetAllOrdersHandler(ctx *gin.Context) {
	res := providers.OrderSrv.ListAllOrders()
	controllers.ReturnJsonStruct(ctx, res)
}

func GetOrderByIdHandler(ctx *gin.Context) {
	pid, err := validators.ValidateOrderByIdReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res := providers.OrderSrv.GetOrdersById(pid)
	controllers.ReturnJsonStruct(ctx, res)
}

func UpdateOrderStatusHandler(ctx *gin.Context) {
	pid, err := validators.ValidateOrderByIdReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res, err := providers.OrderSrv.UpdateOrderById(pid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, res)
}

func CancelOrderHandler(ctx *gin.Context) {
	pid, err := validators.ValidateOrderByIdReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res, err := providers.OrderSrv.CancelOrderById(pid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	controllers.ReturnJsonStruct(ctx, res)
}
