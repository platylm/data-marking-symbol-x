package repository

import (
	"data/data"
	"data/model"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (productRepository ProductRepositoryMySQL) UpdateDataProductBrand(productID int, productBrand string) error {
	markX, _ := data.MarkingSymbolX(productBrand)
	_, err := productRepository.DBConnection.Exec(`UPDATE products SET product_brand = ? WHERE id=?`, markX, productID)
	return err
}

func (productRepository ProductRepositoryMySQL) UpdateAllDataProductBrand() error {
	var result []model.ProductDetail
	err := productRepository.DBConnection.Select(&result, "SELECT id,product_name,product_price,quantity,image_url,product_brand FROM products")
	for _, p := range result {
		err = productRepository.UpdateDataProductBrand(p.ID, p.Brand)
	}
	return err
}
