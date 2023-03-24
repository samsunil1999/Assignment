package validators

import (
	"Assignment/models"
	"Assignment/services/orders"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func ValidateCreateOrderReq(ctx *gin.Context) (models.CreateOrderReq, error) {
	var req models.CreateOrderReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return models.CreateOrderReq{}, err
	}

	if len(req.Inventory) < 1 {
		return models.CreateOrderReq{}, errors.New("atleast one product must be added in the inventory")
	}

	var opts govalidator.Options
	for _, productData := range req.Inventory {
		opts = govalidator.Options{
			Data:  &productData,
			Rules: getRulesForOrderedProductData(),
		}

		v := govalidator.New(opts)
		e := v.ValidateStruct()
		if len(e) > 0 {
			for param, message := range e {
				return models.CreateOrderReq{}, errors.New("param: " + param + ", message:" + message[0])
			}
		}
	}

	return req, nil
}

func ValidateOrderByIdReq(ctx *gin.Context) (string, error) {
	pid := ctx.Param("id")
	_, ok := orders.OrderMap[pid]
	if !ok {
		return "", errors.New("no such order with id: " + pid)
	}

	return pid, nil
}
