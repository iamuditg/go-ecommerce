package models

import (
	"errors"
	"time"
)

type Product struct {
	ID           int       `json:"ID"`
	Name         string    `json:"name"`
	Price        string    `json:"price"`
	Description  string    `json:"description"`
	Rating       float64   `json:"rating"`
	CategoryID   int       `json:"categoryID"`
	SellerID     int       `json:"sellerID"`
	BrandID      int       `json:"brandID"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CategoryName string    `json:"categoryName"`
	SellerName   string    `json:"sellerName"`
	BrandName    string    `json:"brandName"`
}

func (p *Product) Validation() error {
	if p.Name == "" {
		return errors.New("product name is required")
	} else if p.Price == "" {
		return errors.New("product price must be greater then zero")
	}
	return nil
}
