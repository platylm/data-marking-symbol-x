package data_test

import (
	"data/data"
	"data/model"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/go-playground/assert.v1"
)

func Test_MarkSymbolX_Input_Product_Band_Should_Be_XXXX(t *testing.T) {
	expected := "xxxxxx"
	productBand := "Adidas"

	actualProductBand, err := data.MarkingSymbolX(productBand)
	assert.Equal(t, expected, actualProductBand)
	assert.Equal(t, err, nil)
}

func Test_ShuffleProductName_Input_ProductDetail_Should_Be_ProductDetail_Shuffle_Already(t *testing.T) {
	expected := []model.ProductDetail{
		{
			ID:       2,
			Name:     "Adidas Ultra Boots",
			Price:    12.95,
			Quantity: 10,
			Brand:    "CoolKidz",
			Image:    "/43_Piece_dinner_Set.png",
		},
		{
			ID:       3,
			Name:     "43 Piece dinner Set",
			Price:    119.95,
			Quantity: 5,
			Brand:    "Adidas",
			Image:    "/Adidas_Ultra_Boots.png",
		},
	}

	producctDetail := []model.ProductDetail{
		{
			ID:       2,
			Name:     "43 Piece dinner Set",
			Price:    12.95,
			Quantity: 10,
			Brand:    "CoolKidz",
			Image:    "/43_Piece_dinner_Set.png",
		},
		{
			ID:       3,
			Name:     "Adidas Ultra Boots",
			Price:    119.95,
			Quantity: 5,
			Brand:    "Adidas",
			Image:    "/Adidas_Ultra_Boots.png",
		},
	}

	fn := func(int, int) int { return 1 }
	actualShuffleProductName := data.ShuffleProductName(fn, producctDetail)
	assert.Equal(t, expected, actualShuffleProductName)
}
