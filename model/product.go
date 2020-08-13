package model

type ProductDetail struct {
	ID       int     `db:"id"`
	Name     string  `db:"product_name"`
	Price    float64 `db:"product_price"`
	Image    string  `db:"image_url"`
	Quantity int     `db:"quantity"`
	Brand    string  `db:"product_brand"`
}
