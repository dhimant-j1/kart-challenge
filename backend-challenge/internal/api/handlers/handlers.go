package handlers

import (
	"backend-challenge/internal/models"
	"backend-challenge/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product-related API endpoints
type ProductHandler struct {
	productService *service.ProductService
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// ListProducts handles GET /product endpoint
// @Summary     List products
// @Description Get all products available for order
// @Tags        product
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Product
// @Router      /product [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products := h.productService.GetProducts()
	c.JSON(http.StatusOK, products)
}

// GetProduct handles GET /product/{productId} endpoint
// @Summary     Get product by ID
// @Description Returns a single product
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       productId path int true "Product ID"
// @Success     200 {object} models.Product
// @Failure     400 {object} models.APIResponse "Invalid ID supplied"
// @Failure     404 {object} models.APIResponse "Product not found"
// @Router      /product/{productId} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	productID := c.Param("productId")

	// Validate ID (in OpenAPI it's defined as integer)
	_, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid product ID format",
		})
		return
	}

	product, err := h.productService.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Code:    http.StatusNotFound,
			Type:    "error",
			Message: "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// OrderHandler handles order-related API endpoints
type OrderHandler struct {
	orderService *service.OrderService
}

// NewOrderHandler creates a new OrderHandler
func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// PlaceOrder handles POST /order endpoint
// @Summary     Place an order
// @Description Place a new order in the store
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       order body models.OrderRequest true "Order information"
// @Success     200 {object} models.Order
// @Failure     400 {object} models.APIResponse "Invalid input"
// @Failure     401 {object} models.APIResponse "Unauthorized"
// @Failure     403 {object} models.APIResponse "Forbidden"
// @Failure     422 {object} models.APIResponse "Validation exception"
// @Security    ApiKeyAuth
// @Router      /order [post]
func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	// Check API key
	apiKey := c.GetHeader("api_key")
	if apiKey != "apitest" {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Code:    http.StatusUnauthorized,
			Type:    "error",
			Message: "Invalid or missing API key",
		})
		return
	}

	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Validate request
	if len(orderReq.Items) == 0 {
		c.JSON(http.StatusUnprocessableEntity, models.APIResponse{
			Code:    http.StatusUnprocessableEntity,
			Type:    "error",
			Message: "Order must contain at least one item",
		})
		return
	}

	// Process order
	order, err := h.orderService.PlaceOrder(orderReq)
	if err != nil {
		if err.Error() == "invalid coupon: not found in at least two files" ||
			err.Error() == "invalid coupon: must be between 8 and 10 characters" {
			// Coupon specific errors
			c.JSON(http.StatusUnprocessableEntity, models.APIResponse{
				Code:    http.StatusUnprocessableEntity,
				Type:    "error",
				Message: err.Error(),
			})
		} else {
			// Other errors (invalid product, etc)
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, order)
}
