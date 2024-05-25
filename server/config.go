package server

import (
	"database/sql"

	"github.com/Lucifer07/e-wallet/handler"
	"github.com/Lucifer07/e-wallet/repository"
	"github.com/Lucifer07/e-wallet/service"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
)

type RouterOpt struct {
	UserHandler     *handler.UserHandler
	passwordHandler *handler.PasswordTokenHandler
	historyHandler *handler.HistoryHandler
}

func createRouter(db *sql.DB) *gin.Engine {
	helper := util.HelperImpl{}
	transactor := util.NewTransactor(db)

	userRepository := repository.NewUserRepository(db)
	historyRepo:=repository.NewHistoryRepository(db)
	passwordTokenRepo := repository.NewPasswordTokenRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	bankRepo:=repository.NewBankTransactionRepository(db)
	ccRepo:=repository.NewCCTransactionRepository(db)
	paylaterRepo:=repository.NewPayLaterTransactionRepository(db)
	userService := service.NewUserService(userRepository, &helper, walletRepository, transactor)
	historyService:=service.NewHistoryService(historyRepo,walletRepository,transactor,bankRepo,paylaterRepo,ccRepo)
	passwordTokenService := service.NewPasswordTokenService(passwordTokenRepo, transactor, userRepository,&helper)
	userHandler := handler.NewuserHandler(userService)
	passwordTokenHandler := handler.NewPasswordTokenHandler(passwordTokenService)
	historyHandler:=handler.NewHistoryHandler(historyService)
	return NewRouter(RouterOpt{
		UserHandler:     userHandler,
		passwordHandler: passwordTokenHandler,
		historyHandler: historyHandler,
	})
}

func Init(db *sql.DB) *gin.Engine {
	router := createRouter(db)
	return router
}
