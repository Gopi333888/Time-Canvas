package interfaces

import "WatchHive/pkg/utils/models"

type WishListUsecase interface {
	AddToWishList(user_id, product_id int) error
	GetWishList(user_id int) ([]models.WishListData, error)
	RemoveFromWishlist(user_id, product_id int) error
}
