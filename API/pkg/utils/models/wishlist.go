package models

type WishListData struct {
	Product_id int `json:"product_id"`
	Price      float64 `json:"price"`
}
type Wishlist struct {
	UserId int `json:"user_id"`
	ProductId int `json:"product_id"`
}