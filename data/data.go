package data

import (
	"data/model"
	"math/rand"
	"time"
)

type Data interface {
	MarkingSymbolX(productBrand string) (string, error)
	RandomRange(max, min int) int
	ShuffleProductName(randomFn func(int, int) int, productDetail []model.ProductDetail) []model.ProductDetail
}

func MarkingSymbolX(productBrand string) (string, error) {
	var markX string
	for i := 0; i < len(productBrand); i++ {
		markX += "x"
	}
	return markX, nil
}

type Random struct {
	Seed int64
}

func (r Random) Range(max, min int) int {
	rand.Seed(r.Seed)
	rangeNumber := rand.Intn(max-min) + min

	return rangeNumber
}

func RandomRange(max, min int) int {
	secondsSince1970 := time.Now().UTC().UnixNano()
	rand.Seed(secondsSince1970)
	rangeNumber := rand.Intn(max-min) + min

	return rangeNumber
}

func ShuffleProductName(randomFn func(int, int) int, productDetail []model.ProductDetail) []model.ProductDetail {
	for first := range productDetail {
		second := randomFn(len(productDetail), 0)
		productDetail[first].Name, productDetail[second].Name = productDetail[second].Name, productDetail[first].Name
	}
	return productDetail
}
