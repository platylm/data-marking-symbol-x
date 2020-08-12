package data

import (
	"github.com/jmoiron/sqlx"
)

type ProductDetail struct {
	ID       int     `db:"id"`
	Name     string  `db:"product_name"`
	Price    float64 `db:"product_price"`
	Image    string  `db:"image_url"`
	Quantity int     `db:"quantity"`
	Brand    string  `db:"product_brand"`
}

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (productRepository ProductRepositoryMySQL) MarkingSymbolX(productBrand string) (string, error) {
	var markX string
	for i := 0; i < len(productBrand); i++ {
		markX += "x"
	}
	return markX, nil
}

func (productRepository ProductRepositoryMySQL) UpdateDataProductBrand(productID int, productBrand string) error {
	markX, _ := productRepository.MarkingSymbolX(productBrand)
	_, err := productRepository.DBConnection.Exec(`UPDATE products SET product_brand = ? WHERE id=?`, markX, productID)
	return err
}

func (productRepository ProductRepositoryMySQL) UpdateAllDataProductBrand() error {
	var result []ProductDetail
	err := productRepository.DBConnection.Select(&result, "SELECT id,product_name,product_price,quantity,image_url,product_brand FROM products")
	for _, p := range result {
		err = productRepository.UpdateDataProductBrand(p.ID, p.Brand)
	}
	return err
}
