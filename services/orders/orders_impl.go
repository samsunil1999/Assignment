package orders

import (
	"Assignment/constants"
	"Assignment/models"
	"Assignment/services/products"
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderImplementation struct{}

var OrderMap = make(map[string]models.OrderEntity)

func (o OrderImplementation) CreateOrder(req models.CreateOrderReq) (models.CreateOrderRes, error) {
	var premiumProdCount int
	var totalOrderValue float64

	// iterate all products
	for _, product := range req.Inventory {
		p, ok := products.ProductCatalog[product.ProductID]

		// if there is no product with particular product_id
		if !ok {
			return models.CreateOrderRes{}, errors.New("No such product with id: " + product.ProductID)
		}

		// if product_quantity is more than 10 as the limit in max 10
		if product.ProductQuantity > 10 {
			return models.CreateOrderRes{}, errors.New("Cannot place an order with quantity more than 10 (id: " + product.ProductID + ")")
		}

		// if ordered product quantity is more than the product availability
		if p.Availability < product.ProductQuantity {
			return models.CreateOrderRes{}, errors.New("Insufficient availability of product (id: " + product.ProductID + ")")
		}

		if p.Category == constants.ProductCategory.PREMIUM {
			premiumProdCount++
		}

		totalOrderValue += p.Price * float64(product.ProductQuantity)
	}

	// creating uuid for order
	id := "odr_" + uuid.NewString()

	// create order response
	res := models.CreateOrderRes{
		ID:          id,
		Inventory:   req.Inventory,
		OrderAmount: totalOrderValue,
		Status:      "placed",
	}

	// checking eligibility for discount
	// if there is atleast 3 different premium products then give discount of 10%
	if premiumProdCount > 2 {
		res.Discount = true
		res.PayableAmount = res.OrderAmount - (res.OrderAmount / 10)
	} else {
		res.Discount = false
		res.PayableAmount = res.OrderAmount
	}

	// add the order to OrderMaps
	OrderMap[id] = models.OrderEntity{
		Status:        res.Status,
		Product:       req.Inventory,
		Discount:      res.Discount,
		OrderAmount:   res.OrderAmount,
		PayableAmount: res.PayableAmount,
	}

	// update the product catlog ()
	for _, product := range req.Inventory {

		p, ok := products.ProductCatalog[product.ProductID]
		if ok {
			p.Availability -= product.ProductQuantity
		}

		products.ProductCatalog[product.ProductID] = p
	}

	return res, nil
}

func (o OrderImplementation) ListAllOrders() models.ListAllOrdersRes {
	var arrOrder []models.OrderListDetails
	for id, order := range OrderMap {
		orderData := models.OrderListDetails{
			ID:            id,
			Status:        order.Status,
			PayableAmount: order.PayableAmount,
			DispatchDate:  order.DispatchDate,
		}
		arrOrder = append(arrOrder, orderData)
	}

	res := models.ListAllOrdersRes{
		RecordCount: len(OrderMap),
		Data:        arrOrder,
	}

	return res
}

func (o OrderImplementation) GetOrdersById(pid string) models.CreateOrderRes {
	orderData := OrderMap[pid]

	res := models.CreateOrderRes{
		ID:            pid,
		Inventory:     orderData.Product,
		OrderAmount:   orderData.OrderAmount,
		Discount:      orderData.Discount,
		PayableAmount: orderData.PayableAmount,
		Status:        orderData.Status,
		DispatchDate:  orderData.DispatchDate,
	}

	return res
}

func (o OrderImplementation) UpdateOrderById(pid string) (models.UpdateOrderStatus, error) {
	orderData := OrderMap[pid]

	switch orderData.Status {

	case constants.OrderStatus.PLACED:
		orderData.Status = constants.OrderStatus.DISPATCHED
		orderData.DispatchDate = time.Now().Format("02-01-2006")

	case constants.OrderStatus.DISPATCHED:
		orderData.Status = constants.OrderStatus.COMPLETED
		orderData.DispatchDate = ""

	case constants.OrderStatus.COMPLETED, constants.OrderStatus.CANCELLED:
		return models.UpdateOrderStatus{}, errors.New("Cannot change status as the current order is " + string(orderData.Status))
	}

	OrderMap[pid] = orderData

	resp := models.UpdateOrderStatus{
		Status:  "success",
		Message: "Order id: " + pid + ", status updated to " + orderData.Status,
	}

	return resp, nil
}

func (o OrderImplementation) CancelOrderById(pid string) (models.UpdateOrderStatus, error) {
	orderData := OrderMap[pid]

	if orderData.Status == constants.OrderStatus.COMPLETED ||
		orderData.Status == constants.OrderStatus.CANCELLED {
		return models.UpdateOrderStatus{}, errors.New("cannot cancel order, as current order status is" + orderData.Status)
	} else {
		orderData.Status = constants.OrderStatus.CANCELLED
		orderData.DispatchDate = ""
	}

	OrderMap[pid] = orderData
	resp := models.UpdateOrderStatus{
		Status:  "success",
		Message: "Order id: " + pid + ", is " + orderData.Status,
	}
	return resp, nil
}
