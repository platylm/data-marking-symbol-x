package data_test

import (
	"data"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/go-playground/assert.v1"
)

func Test_ProductRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := data.ProductRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("MarkSymbolX_Input_Product_Band_Should_Be_XXXX", func(t *testing.T) {
		expected := "xxxxxx"
		productBand := "Adidas"

		actualProductBand, err := repository.MarkingSymbolX(productBand)
		assert.Equal(t, expected, actualProductBand)
		assert.Equal(t, err, nil)
	})

	t.Run("UpdateDataProductBand_Input_Product_ID_2_And_Product_Band_Adidas_Should_Be_XXXXXX", func(t *testing.T) {
		productID := 3
		productBrand := "Adidas"

		err := repository.UpdateDataProductBrand(productID, productBrand)
		assert.Equal(t, nil, err)
	})

	t.Run("UpdateAllDataProductBrand_Should_Be_XXXXXX", func(t *testing.T) {
		err := repository.UpdateAllDataProductBrand()
		assert.Equal(t, nil, err)
	})
}
