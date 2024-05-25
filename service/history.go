package service

import (
	"context"
	"log"
	"strconv"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/payload"
	"github.com/Lucifer07/e-wallet/repository"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/shopspring/decimal"
)

type HistoryServiceImp struct {
	HistoryRepo  repository.HistoryRepository
	walletRepo   repository.WalletRepository
	transaction  util.Transactor
	bankRepo     repository.BankTransactionRepository
	paylaterRepo repository.PayLaterTransactionRepository
	ccRepo       repository.CCTransactionRepository
}
type HistoryService interface {
	MyTransactions(ctx context.Context, params map[string]string) (*[]entity.HistoryTransaction, error)
	TopupBank(ctx context.Context, dataTopup dto.TopupBankRequest) (*dto.TopupBankResponse, error)
	TopupPayLater(ctx context.Context, dataTopup dto.TopupPaylaterRequest) (*dto.TopupPaylaterResponse, error)
	TopupCreditCard(ctx context.Context, dataTopup dto.TopupCreditCardRequest) (*dto.TopupCreditCardResponse, error)
	Transfer(ctx context.Context, walletTransaction dto.TransferRequest) (*dto.WalletTransactionResponse, error)
}

func NewHistoryService(HistoryRepo repository.HistoryRepository, walletRepo repository.WalletRepository, transaction util.Transactor, bankRepo repository.BankTransactionRepository, paylaterRepo repository.PayLaterTransactionRepository, ccRepo repository.CCTransactionRepository) *HistoryServiceImp {
	return &HistoryServiceImp{
		HistoryRepo:  HistoryRepo,
		walletRepo:   walletRepo,
		transaction:  transaction,
		bankRepo:     bankRepo,
		paylaterRepo: paylaterRepo,
		ccRepo:       ccRepo,
	}
}
func (s *HistoryServiceImp) MyTransactions(ctx context.Context, params map[string]string) (*[]entity.HistoryTransaction, error) {
	var transactions []entity.HistoryTransaction
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		user, err := util.CheckClaim(ctx)
		if err != nil {
			return err
		}
		wallet, err := s.walletRepo.GetWalletData(ctx, user.Id)
		if err != nil {
			return err
		}
		dataTransaction, err := s.HistoryRepo.MyTransactions(ctx, *wallet, params)
		if err != nil {
			return err
		}
		transactions = *dataTransaction
		return nil
	})
	if err != nil {
		return nil, err
	}
	if transactions == nil {
		transactions = []entity.HistoryTransaction{}
	}
	return &transactions, nil
}
func (s *HistoryServiceImp) TopupBank(ctx context.Context, dataTopup dto.TopupBankRequest) (*dto.TopupBankResponse, error) {
	err := checkAmountTopup(dataTopup.Amount)
	if err != nil {
		return nil, err
	}
	var history dto.TopupBankResponse
	err = s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		err := checkAmountTopup(dataTopup.Amount)
		if err != nil {
			return err
		}
		user, err := util.CheckClaim(ctx)
		if err != nil {
			return err
		}
		dataWallet, err := s.walletRepo.GetWalletData(ctx, user.Id)
		if err != nil {
			return err
		}
		bankAccount, err := s.bankRepo.CreateBankAccount(ctx, dataTopup.AccountNumber)
		if err != nil {
			return err
		}
		dataWallet.Balance = dataWallet.Balance.Add(dataTopup.Amount)
		err = s.walletRepo.UpdateBalance(ctx, *dataWallet)
		if err != nil {
			return err
		}
		historyTransaction, err := s.HistoryRepo.CreateHistory(ctx, entity.HistoryTransaction{
			Amount:            dataTopup.Amount,
			TransactionMethod: util.BankTransfer,
			SenderWalletId:    bankAccount.Id,
			RecipientWalletId: dataWallet.Id,
			Description:       dataTopup.Description,
		})
		if err != nil {
			return err
		}

		history = payload.HistoryToBankHistory(*historyTransaction, dataWallet.WalletNumber, bankAccount.AccountNumber)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &history, nil
}
func (s *HistoryServiceImp) TopupPayLater(ctx context.Context, dataTopup dto.TopupPaylaterRequest) (*dto.TopupPaylaterResponse, error) {
	err := checkAmountTopup(dataTopup.Amount)
	if err != nil {
		return nil, err
	}
	var history dto.TopupPaylaterResponse
	err = s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		user, err := util.CheckClaim(ctx)
		if err != nil {
			return err
		}
		dataWallet, err := s.walletRepo.GetWalletData(ctx, user.Id)
		if err != nil {
			return err
		}
		paylatterAccount, err := s.paylaterRepo.CreatePayLaterTransaction(ctx, user.Id)
		if err != nil {
			return err
		}
		dataWallet.Balance = dataWallet.Balance.Add(dataTopup.Amount)
		err = s.walletRepo.UpdateBalance(ctx, *dataWallet)
		if err != nil {
			return err
		}
		historyTransaction, err := s.HistoryRepo.CreateHistory(ctx, entity.HistoryTransaction{
			Amount:            dataTopup.Amount,
			TransactionMethod: util.Paylater,
			SenderWalletId:    paylatterAccount.Id,
			RecipientWalletId: dataWallet.Id,
			Description:       dataTopup.Description,
		})
		if err != nil {
			return err
		}
		history = payload.HistoryToPaylaterHistory(*historyTransaction, dataWallet.WalletNumber, user.Email)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &history, nil
}
func (s *HistoryServiceImp) TopupCreditCard(ctx context.Context, dataTopup dto.TopupCreditCardRequest) (*dto.TopupCreditCardResponse, error) {
	err := checkAmountTopup(dataTopup.Amount)
	if err != nil {
		return nil, err
	}
	var history dto.TopupCreditCardResponse
	err = s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		user, err := util.CheckClaim(ctx)
		if err != nil {
			return err
		}
		dataWallet, err := s.walletRepo.GetWalletData(ctx, user.Id)
		if err != nil {
			return err
		}
		CCAccount, err := s.ccRepo.CreateCCAccount(ctx, dataTopup.CCNumber)
		if err != nil {
			return err
		}
		dataWallet.Balance = dataWallet.Balance.Add(dataTopup.Amount)
		err = s.walletRepo.UpdateBalance(ctx, *dataWallet)
		if err != nil {
			return err
		}
		historyTransaction, err := s.HistoryRepo.CreateHistory(ctx, entity.HistoryTransaction{
			Amount:            dataTopup.Amount,
			TransactionMethod: util.CreditCard,
			SenderWalletId:    CCAccount.Id,
			RecipientWalletId: dataWallet.Id,
			Description:       dataTopup.Description,
		})
		if err != nil {
			return err
		}
		history = payload.HistoryToCCHistory(*historyTransaction, dataWallet.WalletNumber, CCAccount.CCNumber)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &history, nil
}
func (s *HistoryServiceImp) Transfer(ctx context.Context, walletTransaction dto.TransferRequest) (*dto.WalletTransactionResponse, error) {
	err := checkAmountTransfer(walletTransaction.Amount)
	if err != nil {
		return nil, err
	}
	var historyResponse dto.WalletTransactionResponse
	err = s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		user, err := util.CheckClaim(ctx)
		if err != nil {
			return err
		}
		selfWallet, err := s.walletRepo.GetWalletData(ctx, user.Id)
		if err != nil {
			return err
		}
		if selfWallet.Balance.LessThan(walletTransaction.Amount) {
			return util.ErrorBalance
		}
		patrnerWallet, err := s.walletRepo.GetPatnerWallet(ctx, walletTransaction.WalletNumber)
		if err != nil {
			log.Println(err)
			return err
		}
		if patrnerWallet.Id == selfWallet.Id {
			return util.ErrorInvalidTransfer
		}
		selfWallet.Balance = selfWallet.Balance.Sub(walletTransaction.Amount)
		patrnerWallet.Balance = patrnerWallet.Balance.Add(walletTransaction.Amount)
		err = s.walletRepo.UpdateBalance(ctx, *selfWallet)
		if err != nil {
			return err
		}
		err = s.walletRepo.UpdateBalance(ctx, *patrnerWallet)
		if err != nil {
			return err
		}
		senderWalletNumber, _ := strconv.Atoi(selfWallet.WalletNumber)
		patnerWalletNumber, _ := strconv.Atoi(patrnerWallet.WalletNumber)
		history, err := s.HistoryRepo.CreateHistory(ctx, entity.HistoryTransaction{
			TransactionMethod: util.Wallet,
			Amount:            walletTransaction.Amount,
			SenderWalletId:    senderWalletNumber,
			RecipientWalletId: patnerWalletNumber,
			Description:       walletTransaction.Description,
		})
		if err != nil {
			return err
		}
		historyResponse = payload.WalletTransactionToResponse([]entity.Wallet{
			*patrnerWallet,
			*selfWallet,
		}, *history)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &historyResponse, nil
}

func checkAmountTopup(data decimal.Decimal) error {
	if data.LessThan(decimal.NewFromInt(util.MinTopUp)) {
		return util.ErroMinimumTopUp
	}
	if data.GreaterThan(decimal.NewFromInt(util.MaxTopUp)) {
		return util.ErroMaximalTopUp
	}
	return nil
}
func checkAmountTransfer(data decimal.Decimal) error {
	if data.LessThan(decimal.NewFromInt(util.MinTranfer)) {
		return util.ErroMinimumTranfer
	}
	if data.GreaterThan(decimal.NewFromInt(util.MaxTranfer)) {
		return util.ErroMaximalTranfer
	}
	return nil
}
