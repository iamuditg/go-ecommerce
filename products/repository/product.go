package repository

import (
	"database/sql"
	"github.com/iamuditg/models"
	"log"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetProduct() ([]models.Product, error) {
	rows, err := r.DB.Query(`
       SELECT 
             p.id,
             p.name,
             p.price,
             p.description,
             p.rating,
             p.category_id,
             p.seller_id,
             p.brand_id,
             p.image,
             p.created_at,
             p.updated_at,
             c.name AS category_name,
             s.name AS seller_name,
             b.name AS brand_name
           FROM products AS p 
       	   LEFT JOIN categories AS c ON p.category_id = c.id
           LEFT JOIN sellers AS s ON p.seller_id = s.id
       	   LEFT JOIN brands AS b ON p.brand_id = b.id
	`)

	if err != nil {
		log.Printf("failed to query products: %v", err)
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.Rating,
			&product.CategoryID,
			&product.SellerID,
			&product.BrandID,
			&product.Image,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.CategoryName,
			&product.SellerName,
			&product.BrandName,
		)
		if err != nil {
			log.Printf("failed to scan product row: %v", err)
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Printf("error iterating over product rows: %v", err)
		return nil, err
	}
	return products, nil
}
