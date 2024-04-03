package repository

import (
	interfaces "WatchHive/pkg/repository/interface"
	"WatchHive/pkg/utils/models"

	"gorm.io/gorm"
)

type wishListRepository struct {
	DB *gorm.DB
}

func NewWishListRepository(DB *gorm.DB) interfaces.WishListRepository {
	return &wishListRepository{
		DB: DB,
	}
}

func (wr *wishListRepository) AddToWishList(user_id, product_id int) error {
	querry := `insert into wishlists (product_id,user_id) values (?,?)`
	if err := wr.DB.Exec(querry, product_id, user_id).Error; err != nil {
		return err
	}
	return nil

}

func (wr *wishListRepository) GetWishList(user_id int) ([]models.WishListData, error) {

	var wishlist []models.WishListData
	querry := `select product_id from wishlists where user_id = ?`
	if err := wr.DB.Raw(querry, user_id).Scan(&wishlist).Error; err != nil {
		return []models.WishListData{}, err
	}
	return wishlist, nil

}

func (wr *wishListRepository) RemoveFromWishlist(user_id, product_id int) error {
	querry := `delete from wishlists where user_id = ? and product_id = ?`
	if err := wr.DB.Exec(querry, user_id, product_id).Error; err != nil {
		return err
	}
	return nil

}
func (wr *wishListRepository) IsExistOnWishlist(user_id, product_id int) (bool, error) {
	var count int
	querry := `select count(*) from wishlists where user_id = ? and product_id = ?`
	if err := wr.DB.Raw(querry, user_id, product_id).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil

}
