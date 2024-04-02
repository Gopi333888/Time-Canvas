package usecase

import (
	repo "WatchHive/pkg/repository/interface"
	interfaces "WatchHive/pkg/usecase/interface"
	"WatchHive/pkg/utils/errmsg"
	"WatchHive/pkg/utils/models"
	"errors"
)

type wishListUsecase struct {
	wishlist repo.WishListRepository
}

func NewWishListUsecase(wl repo.WishListRepository) interfaces.WishListUsecase {
	return &wishListUsecase{
		wishlist: wl,
	}
}

func (wu *wishListUsecase) AddToWishList(user_id, product_id int) error {
	err := wu.wishlist.AddToWishList(user_id, product_id)
	if err != nil {
		return errors.New(errmsg.ErrWriteDB)
	}
	return nil
}

func (wu *wishListUsecase) GetWishList(user_id int) ([]models.WishListData, error) {
	wishlist, err := wu.wishlist.GetWishList(user_id)
	if err != nil {
		return []models.WishListData{}, errors.New(errmsg.ErrGetData)
	}
	return wishlist, nil
}

func (wu *wishListUsecase) RemoveFromWishlist(user_id, product_id int) error {
	err := wu.wishlist.RemoveFromWishlist(user_id, product_id)
	if err != nil {
		return errors.New(errmsg.ErrInternal)
	}
	return nil
}
