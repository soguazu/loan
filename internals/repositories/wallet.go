package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type walletRepository struct {
	db *gorm.DB
}

// NewWalletRepository creates a new instance wallet repository
func NewWalletRepository(db *gorm.DB) ports.IWalletRepository {
	return &walletRepository{
		db: db,
	}
}

func (w *walletRepository) GetByID(id string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	if err := w.db.Where("id = ?", id).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (w *walletRepository) GetByIDForUpdate(id string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	if err := w.db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	}).Where("id = ?", id).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (w *walletRepository) GetByCompany(id string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	if err := w.db.Where("company = ?", id).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (w *walletRepository) GetBy(filter interface{}) ([]domain.Wallet, error) {
	var wallet []domain.Wallet
	if err := w.db.Model(&domain.Wallet{}).Find(&wallet, filter).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (w *walletRepository) Persist(wallet *domain.Wallet) error {
	if wallet.ID.String() != "" {
		if err := w.db.Save(wallet).Error; err != nil {
			return err
		}
		return nil
	}
	if err := w.db.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}

func (w *walletRepository) Delete(id string) error {
	if err := w.db.Where("id = ?", id).Delete(&domain.Wallet{}).Error; err != nil {
		return err
	}
	return nil
}

func (w *walletRepository) DeleteAll() error {
	if err := w.db.Exec("DELETE FROM wallets").Error; err != nil {
		return err
	}
	return nil
}

func (w *walletRepository) WithTx(tx *gorm.DB) ports.IWalletRepository {
	return NewWalletRepository(tx)
}
