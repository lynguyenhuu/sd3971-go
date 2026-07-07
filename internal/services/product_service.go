package services

import (
	"fmt"

	"sd3971-go/internal/models"
	"sd3971-go/internal/repositories"
)

// ProductService handles business logic for products
type ProductService struct {
	repo *repositories.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetAllProducts retrieves all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}
	return products, nil
}

// GetProductByID retrieves a product by ID
func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid product ID")
	}

	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(req *models.CreateProductRequest) (*models.Product, error) {
	// Validate input
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	createdProduct, err := s.repo.Create(product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return createdProduct, nil
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(id int, req *models.UpdateProductRequest) (*models.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid product ID")
	}

	// Retrieve existing product
	existingProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Apply updates
	if req.Name != nil {
		existingProduct.Name = *req.Name
	}
	if req.Description != nil {
		existingProduct.Description = *req.Description
	}
	if req.Price != nil {
		existingProduct.Price = *req.Price
	}
	if req.Quantity != nil {
		existingProduct.Quantity = *req.Quantity
	}

	// Validate updated product
	if err := validateProduct(existingProduct); err != nil {
		return nil, err
	}

	updatedProduct, err := s.repo.Update(id, existingProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return updatedProduct, nil
}

// DeleteProduct deletes a product
func (s *ProductService) DeleteProduct(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid product ID")
	}

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// validateCreateRequest validates the create product request
func validateCreateRequest(req *models.CreateProductRequest) error {
	if req.Name == "" {
		return fmt.Errorf("product name is required")
	}
	if req.Description == "" {
		return fmt.Errorf("product description is required")
	}
	if req.Price <= 0 {
		return fmt.Errorf("product price must be greater than 0")
	}
	if req.Quantity < 0 {
		return fmt.Errorf("product quantity cannot be negative")
	}
	return nil
}

// validateProduct validates a product
func validateProduct(product *models.Product) error {
	if product.Name == "" {
		return fmt.Errorf("product name is required")
	}
	if product.Description == "" {
		return fmt.Errorf("product description is required")
	}
	if product.Price <= 0 {
		return fmt.Errorf("product price must be greater than 0")
	}
	if product.Quantity < 0 {
		return fmt.Errorf("product quantity cannot be negative")
	}
	return nil
}
