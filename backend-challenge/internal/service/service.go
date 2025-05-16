package service

import (
	"backend-challenge/internal/models"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// ProductService handles product-related business logic
type ProductService struct {
	products []models.Product
}

// NewProductService creates a new ProductService
func NewProductService() *ProductService {
	return &ProductService{
		products: models.MockProducts(),
	}
}

// GetProducts returns all available products
func (s *ProductService) GetProducts() []models.Product {
	return s.products
}

// GetProductByID returns a product by its ID or error if not found
func (s *ProductService) GetProductByID(id string) (models.Product, error) {
	for _, product := range s.products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

// OrderService handles order-related business logic
type OrderService struct {
	productService *ProductService
	couponCache    map[string]bool
	cacheMutex     sync.RWMutex
}

// NewOrderService creates a new OrderService
func NewOrderService(productService *ProductService) *OrderService {
	return &OrderService{
		productService: productService,
		couponCache:    make(map[string]bool),
	}
}

// PlaceOrder processes an order and returns the order details
func (s *OrderService) PlaceOrder(req models.OrderRequest) (models.Order, error) {
	// Validate items
	if len(req.Items) == 0 {
		return models.Order{}, errors.New("order must contain at least one item")
	}

	// Create empty order
	order := models.Order{
		ID:        "order-" + strings.Replace(generateUUID(), "-", "", -1)[:16],
		Items:     req.Items,
		Products:  make([]models.Product, 0, len(req.Items)),
		Discounts: 0,
	}

	// Calculate total and add products
	var total float64
	for _, item := range req.Items {
		product, err := s.productService.GetProductByID(item.ProductID)
		if err != nil {
			return models.Order{}, fmt.Errorf("invalid product ID: %s", item.ProductID)
		}
		order.Products = append(order.Products, product)
		total += product.Price * float64(item.Quantity)
	}
	order.Total = total

	// Apply coupon if provided
	if req.CouponCode != "" {
		isValid, err := s.ValidateCoupon(req.CouponCode)
		if err != nil {
			return models.Order{}, err
		}
		if isValid {
			// Apply 10% discount
			discount := order.Total * 0.1
			order.Discounts = discount
			order.Total -= discount
		}
	}

	return order, nil
}

// ValidateCoupon checks if a coupon code is valid
func (s *OrderService) ValidateCoupon(code string) (bool, error) {
	// Check cache first
	s.cacheMutex.RLock()
	valid, found := s.couponCache[code]
	s.cacheMutex.RUnlock()

	if found {
		return valid, nil
	}

	// Validate coupon length (8-10 characters)
	codeLen := len(code)
	if codeLen < 8 || codeLen > 10 {
		return false, errors.New("invalid coupon: must be between 8 and 10 characters")
	}

	// Check coupon existence in at least two files
	dataDir := filepath.Join("data")
	files := []string{
		filepath.Join(dataDir, "couponbase1.gz"),
		filepath.Join(dataDir, "couponbase2.gz"),
		filepath.Join(dataDir, "couponbase3.gz"),
	}

	occurrences := 0
	for _, file := range files {
		exists, err := couponExistsInFile(file, code)
		if err != nil {
			return false, fmt.Errorf("error validating coupon in %s: %v", file, err)
		}
		if exists {
			occurrences++
			if occurrences >= 2 {
				// Coupon is valid, update cache and return
				s.cacheMutex.Lock()
				s.couponCache[code] = true
				s.cacheMutex.Unlock()
				return true, nil
			}
		}
	}

	// Cache invalid coupon
	s.cacheMutex.Lock()
	s.couponCache[code] = false
	s.cacheMutex.Unlock()

	return false, errors.New("invalid coupon: not found in at least two files")
}

// couponExistsInFile checks if a coupon exists in a gzipped file
func couponExistsInFile(filepath string, coupon string) (bool, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return false, err
	}
	defer gzr.Close()

	// Read the uncompressed file contents and look for the coupon
	contents, err := io.ReadAll(gzr)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(contents), coupon), nil
}

// Helper function to generate a UUID for order IDs
func generateUUID() string {
	return uuid.New().String()
}
