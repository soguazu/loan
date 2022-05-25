package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance transaction repository
func NewTransactionRepository(db *gorm.DB) ports.ITransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) GetByID(id string) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := t.db.Where("id = ?", id).
		Preload("Company").
		Preload("Customer").
		Preload("Fee").
		Preload("ExpenseCategory ").
		First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *transactionRepository) GetTransactionByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error) {
	var transactions []domain.Transaction
	if err := t.db.Scopes(utils.Paginate(transactions, pagination, t.db)).
		Where("Company = ?", id).
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	pagination.Rows = transactions
	return pagination, nil
}

func (t *transactionRepository) GetTransactionByCardID(id string, pagination *utils.Pagination) (*utils.Pagination, error) {
	var transactions []domain.Transaction
	if err := t.db.Scopes(utils.Paginate(transactions, pagination, t.db)).
		Where("Card = ?", id).
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	pagination.Rows = transactions
	return pagination, nil
}

func (t *transactionRepository) GetBy(filter interface{}) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := t.db.Model(&domain.Transaction{}).Find(&transactions, filter).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var transactions []domain.Transaction
	if err := t.db.Scopes(utils.Paginate(transactions, pagination, t.db)).Find(&transactions).Error; err != nil {
		return nil, err
	}
	pagination.Rows = transactions
	return pagination, nil
}

func (t *transactionRepository) Persist(transaction *domain.Transaction) error {
	if transaction.ID.String() != "" {
		if err := t.db.Save(transaction).Error; err != nil {
			return err
		}
		return nil
	}
	if err := t.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) Delete(id string) error {
	if err := t.db.Where("id = ?", id).Delete(&domain.Transaction{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) DeleteAll() error {
	if err := t.db.Exec("DELETE FROM transactions").Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) WithTx(tx *gorm.DB) ports.ITransactionRepository {
	return NewTransactionRepository(tx)
}
