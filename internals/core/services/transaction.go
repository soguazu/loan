package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type transactionService struct {
	TransactionRepository ports.ITransactionRepository
	logger                *log.Logger
}

// NewTransactionService function create a new instance for service
func NewTransactionService(cr ports.ITransactionRepository, l *log.Logger) ports.ITransactionService {
	return &transactionService{
		TransactionRepository: cr,
		logger:                l,
	}
}

func (ts *transactionService) GetTransactionByID(id string) (*domain.Transaction, error) {
	address, err := ts.TransactionRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (ts *transactionService) GetTransactionByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error) {
	transactions, err := ts.TransactionRepository.GetTransactionByCompanyID(id, pagination)
	if err != nil {
		ts.logger.Error(err)
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) GetTransactionByCardID(id string, pagination *utils.Pagination) (*utils.Pagination, error) {
	transactions, err := ts.TransactionRepository.GetTransactionByCompanyID(id, pagination)
	if err != nil {
		ts.logger.Error(err)
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) GetAllTransaction(pagination *utils.Pagination) (*utils.Pagination, error) {
	transactions, err := ts.TransactionRepository.Get(pagination)
	if err != nil {
		ts.logger.Error(err)
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) CreateTransaction(transaction *domain.Transaction) error {
	err := ts.TransactionRepository.Persist(transaction)
	if err != nil {
		ts.logger.Error(err)
		return err
	}
	return nil
}

func (ts *transactionService) UpdateTransaction(id string, body common.UpdateTransactionRequest) (*domain.Transaction, error) {
	transaction, err := ts.TransactionRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if body.Receipt != nil {
		transaction.Receipt = *body.Receipt
	}

	if body.ExpenseCategory != nil {
		transaction.ExpenseCategory, err = uuid.FromString(*body.ExpenseCategory)
		if err != nil {
			return nil, err
		}
	}

	err = ts.TransactionRepository.Persist(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (ts *transactionService) DeleteTransaction(id string) error {
	err := ts.TransactionRepository.Delete(id)
	if err != nil {
		ts.logger.Error(err)
		return err
	}
	return nil
}

func (ts *transactionService) LockTransaction(id string) (*domain.Transaction, error) {
	transaction, err := ts.TransactionRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	if transaction.Lock == false {
		transaction.Lock = true
	} else {
		transaction.Lock = false
	}

	err = ts.TransactionRepository.Persist(transaction)

	if err != nil {
		ts.logger.Error(err)
		return nil, err
	}
	return transaction, nil
}
