package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	log "github.com/sirupsen/logrus"
)

type walletService struct {
	WalletRepository ports.IWalletRepository
	logger           *log.Logger
}

// NewWalletService function create a new instance for service
func NewWalletService(cr ports.IWalletRepository, l *log.Logger) ports.IWalletService {
	return &walletService{
		WalletRepository: cr,
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
	if body.Payment != nil {
		wallet.Payment = *body.Payment
	}

	err = ws.WalletRepository.Persist(wallet)

	if err != nil {
		ws.logger.Error(err)
		return nil, err
	}
	return wallet, nil
}
