package di

// import (
// 	http "WatchHive/pkg/api"
// 	"WatchHive/pkg/api/handler"
// 	"WatchHive/pkg/config"
// 	"WatchHive/pkg/db"
// 	"WatchHive/pkg/helper"
// 	"WatchHive/pkg/repository"
// 	"WatchHive/pkg/usecase"

// 	"github.com/google/wire"
// )

// func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
// 	wire.Build(

// 		db.ConnectDatabase,

// 		repository.NewUserRepository,
// 		repository.NewAdminRepository,
// 		repository.NewOtpRepository,
// 		repository.NewCategoryRepository,
// 		repository.NewProductRepository,
// 		repository.NewCartRepository,
// 		repository.NewOrderRepository,
// 		repository.NewPaymentRepository,
// 		repository.NewWalletRepository,
// 		repository.NewOfferRepository,
// 		repository.NewCouponRepository,
// 		repository.NewWishListRepository,

// 		usecase.NewUserUseCase,
// 		usecase.NewAdminUseCase,
// 		usecase.NewOtpUseCase,
// 		usecase.NewCategoryUseCase,
// 		usecase.NewProductUseCase,
// 		usecase.NewCartUseCase,
// 		usecase.NewOrderUseCase,
// 		usecase.NewPaymentUseCase,
// 		usecase.NewWalletUsecase,
// 		usecase.NewOferUsecase,
// 		usecase.NewCouponUsecase,
// 		usecase.NewWishListUsecase,

// 		handler.NewUserHandler,
// 		handler.NewAdminHandler,
// 		handler.NewOtpHandler,
// 		handler.NewCategoryHandler,
// 		handler.NewProductHandler,
// 		handler.NewCartHandler,
// 		handler.NewOrderHandler,
// 		handler.NewPaymentHandler,
// 		handler.NewWalletHandler,
// 		handler.NewOfferHandler,
// 		handler.NewCouponHandler,
// 		handler.NewWishListHandler,

// 		helper.NewHelper,

// 		http.NewServerHTTP,
// 	)

// 	return &http.ServerHTTP{}, nil
// }
