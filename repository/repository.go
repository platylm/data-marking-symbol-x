package repository

import (
	"data/model"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetAllProduct() ([]model.ProductDetail, error)
	UpdateDataProductBrand(productID int, productBrand string) error
	UpdateResultShuffleProductName(productID int, productName string) error
}

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (productRepository ProductRepositoryMySQL) GetAllProduct() ([]model.ProductDetail, error) {
	var result []model.ProductDetail
	err := productRepository.DBConnection.Select(&result, "SELECT id,product_name,product_price,quantity,image_url,product_brand FROM products")
	return result, err
}

func (productRepository ProductRepositoryMySQL) UpdateDataProductBrand(productID int, productBrand string) error {
	_, err := productRepository.DBConnection.Exec(`UPDATE products SET product_brand = ? WHERE id=?`, productBrand, productID)
	return err
}

func (productRepository ProductRepositoryMySQL) UpdateResultShuffleProductName(productID int, productName string) error {
	_, err := productRepository.DBConnection.Exec(`UPDATE products SET product_name = ? WHERE id=?`, productName, productID)
	return err
}
