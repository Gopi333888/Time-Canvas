package domain

type Wishlist struct {
	ID        uint  `json:"id" gorm:"primaryKey"`
	ProductId int   `json:"product_id"`
	UserId    uint  `json:"user_id"`
	User      Users `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}
