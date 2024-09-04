package service

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"

	"github.com/google/uuid"
)

type orderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) ports.OrderService {
	return &orderService{repo: repo}
}

// CreateOrder implements ports.OrderService.
func (o *orderService) CreateOrder(request *request.NewOrderReuqest) (*response.OrderResponse, error) {

	// find cart items by user id
	cartItems, err := o.repo.FindCartItemByUserId(request.UserID)
	if err != nil {
		return nil, err
	}

	productIds := make([]uint, 0, len(cartItems))
	for _, cartItem := range cartItems {
		productIds = append(productIds, cartItem.ProductID)
	}

	products, err := o.repo.FindProductByIds(productIds)
	if err != nil {
		return nil, err
	}
	// create a map of product id to product
	productMap := make(map[uint]domain.Product)

	for _, product := range products {
		productMap[product.ID] = product
	}

	// Create a new order
	order := &domain.Order{
		OrderNumber: uuid.New(),
		UserID:      request.UserID,
		Status:      domain.Pending,
		TotalPay:    request.TotalPay,
	}

	if err := o.repo.CreateOrder(order); err != nil {
		return nil, err
	}

	// create order items
	orderItems := make([]domain.OrderItem, 0, len(cartItems))
	for _, cartItem := range cartItems {
		if product, ok := productMap[cartItem.ProductID]; ok {
			orderItems = append(orderItems, domain.OrderItem{
				OrderNumber: order.OrderNumber,
				UserID:      order.UserID,
				ProductID:   cartItem.ProductID,
				Quantity:    cartItem.Quantity,
				Price:       product.Price,
			})
		}
	}

	// create order items
	if err := o.repo.CreateOrderItems(orderItems); err != nil {
		return nil, err
	}

	// delete cart items
	if err := o.repo.DeleteCartItemByUserId(request.UserID); err != nil {
		return nil, err
	}

	response := &response.OrderResponse{
		ID:          order.ID,
		OrderNumber: order.OrderNumber,
		Status:      order.Status,
		TotalPay:    order.TotalPay,
	}

	return response, nil
}

// GetOrderByStatus implements ports.OrderService.
func (o *orderService) GetOrderByStatus(request *request.FindOrderByStatusRequest) ([]response.OrderHistoryResponse, error) {

	orders, err := o.repo.FindOrderByUserIdAndStatus(request.UserID, request.Status)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return []response.OrderHistoryResponse{}, nil
	}

	responses := make([]response.OrderHistoryResponse, 0, len(orders))
	for _, order := range orders {
		orderItems, err := o.repo.FindOrderItemByOrderNumber(order.OrderNumber)
		if err != nil {
			return nil, err
		}
		productIds := make([]uint, 0, len(orderItems))

		for _, orderItem := range orderItems {
			productIds = append(productIds, orderItem.ProductID)
		}

		products, err := o.repo.FindProductByIds(productIds)
		if err != nil {
			return nil, err
		}

		productMap, err := getProductMap(products)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		orderItemsResponse := make([]response.OrderItemResponse, 0, len(orderItems))
		for _, orderItem := range orderItems {
			if product, ok := productMap[orderItem.ProductID]; ok {
				orderItemsResponse = append(orderItemsResponse, response.OrderItemResponse{
					ProductName: product.Name,
					Quantity:    orderItem.Quantity,
					Price:       orderItem.Price,
				})
			}
		}

		responses = append(responses, response.OrderHistoryResponse{
			ID:          order.ID,
			OrderNumber: order.OrderNumber,
			Status:      order.Status,
			TotalPay:    order.TotalPay,
			OrderItems:  orderItemsResponse,
		})
	}
	return responses, nil
}

// GetOrderHistory implements ports.OrderService.
func (o *orderService) GetOrderHistory(user *domain.User) ([]response.OrderHistoryResponse, error) {

	orders, err := o.repo.FindOrderByUserId(user.ID)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return []response.OrderHistoryResponse{}, nil
	}

	responses := make([]response.OrderHistoryResponse, 0, len(orders))
	for _, order := range orders {
		orderItems, err := o.repo.FindOrderItemByOrderNumber(order.OrderNumber)
		if err != nil {
			return nil, err
		}
		productIds := make([]uint, 0, len(orderItems))

		for _, orderItem := range orderItems {
			productIds = append(productIds, orderItem.ProductID)
		}

		products, err := o.repo.FindProductByIds(productIds)
		if err != nil {
			return nil, err
		}

		productMap, err := getProductMap(products)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		orderItemsResponse := make([]response.OrderItemResponse, 0, len(orderItems))
		for _, orderItem := range orderItems {
			if product, ok := productMap[orderItem.ProductID]; ok {
				orderItemsResponse = append(orderItemsResponse, response.OrderItemResponse{
					ProductName: product.Name,
					Quantity:    orderItem.Quantity,
					Price:       orderItem.Price,
				})
			}
		}

		responses = append(responses, response.OrderHistoryResponse{
			ID:          order.ID,
			OrderNumber: order.OrderNumber,
			Status:      order.Status,
			TotalPay:    order.TotalPay,
			OrderItems:  orderItemsResponse,
		})
	}
	return responses, nil
}

// CancelOrder implements ports.OrderService.
func (o *orderService) CancelOrder(orderId uint) error {
	order, err := o.repo.FindOrderByID(orderId)
	if err != nil {
		return err
	}

	// return product quantity
	orderItems, err := o.repo.FindOrderItemByOrderNumber(order.OrderNumber)
	if err != nil {
		return err
	}

	for _, orderItem := range orderItems {
		product, err := o.repo.FindProudctById(orderItem.ProductID)
		if err != nil {
			return err
		}
		product.Quantity = product.Quantity + orderItem.Quantity

		if err := o.repo.UpdateProductQuantityById(product); err != nil {
			return err
		}
	}

	order.Status = domain.Cancel

	// update order status
	if err := o.repo.UpdateOrder(order); err != nil {
		return err
	}

	return nil
}
