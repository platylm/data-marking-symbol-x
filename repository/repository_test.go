package repository_test

import (
	"data/repository"
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
	repository := repository.ProductRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("UpdateDataProductBand_Input_Product_ID_2_And_Product_Band_Adidas_Should_Be_XXXXXX", func(t *testing.T) {
		productID := 3
		productBrand := "Adidas"

		err := repository.UpdateDataProductBrand(productID, productBrand)
		assert.Equal(t, nil, err)
	})

	t.Run("UpdateResultShuffleProductName_Should_Be_Updated", func(t *testing.T) {
		productID := 3
		productName := "Adidas Ultra Boots"

		err := repository.UpdateResultShuffleProductName(productID, productName)
		assert.Equal(t, nil, err)
	})
}
