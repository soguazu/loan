package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	tx "core_business/pkg/unit_of_work"
	"core_business/pkg/utils"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type walletService struct {
	WalletRepository ports.IWalletRepository
	DB               *gorm.DB
	logger           *log.Logger
}

// NewWalletService function create a new instance for service
func NewWalletService(cr ports.IWalletRepository, db *gorm.DB, l *log.Logger) ports.IWalletService {
	return &walletService{
		WalletRepository: cr,
		DB:               db,
		logger:           l,
	}
}

func (ws *walletService) GetWalletByID(id string) (*domain.Wallet, error) {
	wallet, err := ws.WalletRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (ws *walletService) CreateWallet(wallet *domain.Wallet) error {
	err := ws.WalletRepository.Persist(wallet)
	if err != nil {
		ws.logger.Error(err)
		return err
	}
	return nil
}

func (ws *walletService) DeleteWallet(id string) error {
	err := ws.WalletRepository.Delete(id)
	if err != nil {
		ws.logger.Error(err)
		return err
	}
	return nil
}

func (ws *walletService) UpdateWallet(id string, body common.UpdateWalletRequest) (*domain.Wallet, error) {

	wallet, err := ws.WalletRepository.GetByID(id)
	if err != nil {
		ws.logger.Error(err)
		return nil, err
	}

	if body.CreditLimit != nil {
		wallet.CreditLimit = *body.CreditLimit
	}

	err = ws.WalletRepository.Persist(wallet)

	if err != nil {
		ws.logger.Error(err)
		return nil, err
	}

	return wallet, nil
}

func (ws *walletService) UpdateBalance(id string, body common.UpdateWalletRequest) (*domain.Wallet, error) {
	uw := tx.NewGormUnitOfWork(ws.DB)
	txx, err := uw.Begin()

	defer func() {
		if err != nil {
			txx.Rollback()
		}
	}()

	wallet, err := ws.WalletRepository.WithTx(txx).GetByIDForUpdate(id)
	if err != nil {
		ws.logger.Error(err)
		return nil, err
	}

	amountInKobo := utils.ToMinorUnit(*body.Payment)

	if *body.Type == "debit" {
		if wallet.AvailableCredit > amountInKobo {
			wallet.CurrentSpending += amountInKobo
			wallet.TotalBalance = wallet.CurrentSpending + (wallet.PreviousBalance - wallet.CashBackPayment)
			wallet.AvailableCredit = wallet.CreditLimit - wallet.TotalBalance
		} else {
			return nil, errors.New("insufficient available credit")
		}
	} else if *body.Type == "credit" {
		wallet.CashBackPayment += amountInKobo
		wallet.TotalBalance = wallet.CurrentSpending + (wallet.PreviousBalance - wallet.CashBackPayment)
		wallet.AvailableCredit = wallet.CreditLimit - wallet.TotalBalance
	}

	err = ws.WalletRepository.WithTx(txx).Persist(wallet)

	if err != nil {
		ws.logger.Error(err)
		return nil, err
	}

	uw.Commit()
	return wallet, nil
}
