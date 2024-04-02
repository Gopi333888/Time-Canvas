package handler

import (
	interfaces "WatchHive/pkg/usecase/interface"
	"WatchHive/pkg/utils/errmsg"
	"WatchHive/pkg/utils/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishListhandler struct {
	WishListUsecase interfaces.WishListUsecase
}

func NewWishListHandler(usecase interfaces.WishListUsecase) *WishListhandler {
	return &WishListhandler{
		WishListUsecase: usecase,
	}
}

// AddToWishList adds a product to the user's wishlist.
// @Summary Add to Wishlist
// @Description Add a product to the user's wishlist.
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Tags User Wishlist Management
// @Param product_id query int true "Product ID" Format(int64)
// @Success 200 {object} response.Response "Successfully added to wishlist"
// @Failure 400 {object} response.Response "Invalid request format or constraints not satisfied"
// @Router /wishlist [post]
func (wh *WishListhandler) AddToWishList(c *gin.Context) {

	id, ok := c.Get("id")
	if !ok {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgAddErr, nil, errors.New(errmsg.ErrGetUserId))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	pId := c.Query("product_id")
	pID, err := strconv.Atoi(pId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgAddErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	iD := id.(int)

	err = wh.WishListUsecase.AddToWishList(iD, pID)
	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgAddErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, errmsg.MsgAddSuccess, nil, nil)
	c.JSON(http.StatusOK, successResp)
}

// GetWishlist retrieves the user's wishlist.
// @Summary Retrieve Wishlist
// @Description Retrieves the user's wishlist.
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Tags User Wishlist Management
// @Success 200 {object} response.Response "Successfully retrieved wishlist"
// @Failure 400 {object} response.Response "Invalid request format or constraints not satisfied"
// @Router /wishlist [get]
func (wh *WishListhandler) GetWishlist(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgAddErr, nil, errors.New(errmsg.ErrGetUserId))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	iD := id.(int)
	wishlistData, err := wh.WishListUsecase.GetWishList(iD)

	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, errmsg.MsgAddSuccess, wishlistData, nil)
	c.JSON(http.StatusOK, successResp)
}

// RemoveFromWishlist removes a product from the user's wishlist.
// @Summary Remove from Wishlist
// @Description Removes a product from the user's wishlist.
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Tags User Wishlist Management
// @Param product_id query int true "Product ID" Format(int64)
// @Success 200 {object} response.Response "Successfully removed from wishlist"
// @Failure 400 {object} response.Response "Invalid request format or constraints not satisfied"
// @Router /wishlist [delete]
func (wh *WishListhandler) RemoveFromWishlist(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgRemoveFailed, nil, errors.New(errmsg.ErrGetUserId))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	pId := c.Query("product_id")
	pID, err := strconv.Atoi(pId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgRemoveFailed, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	iD := id.(int)
	err = wh.WishListUsecase.RemoveFromWishlist(iD, pID)

	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, errmsg.MsgRemoveFailed, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, errmsg.MsgAddSuccess, nil, nil)
	c.JSON(http.StatusOK, successResp)
}
