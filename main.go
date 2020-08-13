package main

import (
	"data/data"
	"data/repository"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ProductRepository repository.ProductRepository
}

func main() {
	dbConnecton := "sealteam:sckshuhari@(localhost:3306)/toy"
	if os.Getenv("DBCONNECTION") != "" {
		dbConnecton = os.Getenv("DBCONNECTION")
	}
	connection, err := sqlx.Connect("mysql", dbConnecton)
	if err != nil {
		log.Fatalln("cannot connect to database", err)
	}

	productRepository := repository.ProductRepositoryMySQL{
		DBConnection: connection,
	}

	getAllProduct, err := productRepository.GetAllProduct()

	for _, p := range getAllProduct {
		markSymbolX, err := data.MarkingSymbolX(p.Brand)
		if err != nil {
			log.Fatalln("cannot stamp symbol X", err)
		}
		err = productRepository.UpdateDataProductBrand(p.ID, markSymbolX)
	}

	fn := func(int, int) int { return len(getAllProduct) - 1 }
	productDetail := data.ShuffleProductName(fn, getAllProduct)

	for _, p := range productDetail {
		err = productRepository.UpdateResultShuffleProductName(p.ID, p.Name)
	}
}
