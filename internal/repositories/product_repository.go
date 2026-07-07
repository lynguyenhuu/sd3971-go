package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"sd3971-go/internal/models"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll retrieves all products from the database
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := `
		SELECT id, name, description, price, quantity, created_at, updated_at
		FROM products
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID, &product.Name, &product.Description,
			&product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

// GetByID retrieves a single product by ID
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT id, name, description, price, quantity, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	var product models.Product
	err := r.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Description,
		&product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query product: %w", err)
	}

	return &product, nil
}

// Create inserts a new product into the database
func (r *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	query := `
		INSERT INTO products (name, description, price, quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		product.Name, product.Description, product.Price, product.Quantity, now, now,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

// Update modifies an existing product in the database
func (r *ProductRepository) Update(id int, product *models.Product) (*models.Product, error) {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, quantity = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, name, description, price, quantity, created_at, updated_at
	`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		product.Name, product.Description, product.Price, product.Quantity, now, id,
	).Scan(
		&product.ID, &product.Name, &product.Description,
		&product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return product, nil
}

// Delete removes a product from the database
func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}
