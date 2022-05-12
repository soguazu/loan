package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"strings"
)

type transactionService struct {
	TransactionRepository ports.ITransactionRepository
	CustomerRepository    ports.ICustomerRepository
	WalletRepository      ports.IWalletRepository
	WalletService         ports.IWalletService
	FeeRepository         ports.IFeeRepository
	CompanyRepository     ports.ICompanyRepository
	CardRepository        ports.ICardRepository
	logger                *log.Logger
}

// NewTransactionService function create a new instance for service
func NewTransactionService(tr ports.ITransactionRepository,
	cr ports.ICustomerRepository, wr ports.IWalletRepository,
	fr ports.IFeeRepository, cmr ports.ICompanyRepository,
	cdr ports.ICardRepository,
	ws ports.IWalletService, l *log.Logger) ports.ITransactionService {
	return &transactionService{
		TransactionRepository: tr,
		CustomerRepository:    cr,
		WalletRepository:      wr,
		WalletService:         ws,
		FeeRepository:         fr,
		CompanyRepository:     cmr,
		CardRepository:        cdr,
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

func (ts *transactionService) CreateTransaction(body *common.CreateTransactionRequest) error {

	var (
		chargeIdentify common.PricingIdentifier
	)

	payload := body.Data.Object
	customerEntity := domain.Customer{
		PartnerCustomerID: payload.Customer.Id,
	}

	customer, err := ts.CustomerRepository.GetBy(customerEntity)
	if err != nil {
		return err
	}

	channel := strings.ToLower(payload.TransactionMetadata.Channel)

	if channel == "web" || channel == "pos" {
		chargeIdentify = common.CardTransactionBOTH
	} else if channel == "atm" {
		chargeIdentify = common.CardTransactionATM
	} else if channel == "" {
		return errors.New("invalid channel")
	}

	identifier, err := ts.FeeRepository.GetByIdentifier(string(chargeIdentify))

	if err != nil {
		return err
	}

	fee := identifier.Fee / 100 * float64(payload.PendingRequest.Amount)

	totalAmount := fee + float64(payload.PendingRequest.Amount)

	wallet, err := ts.WalletRepository.GetByCompany(customer.Company.String())

	if err != nil {
		return err
	}

	debitWallet := common.UpdateWalletRequest{
		CreditLimit:     &wallet.CreditLimit,
		PreviousBalance: &wallet.PreviousBalance,
		CurrentSpending: &wallet.CurrentSpending,
		Payment:         &totalAmount,
	}

	_, err = ts.WalletService.UpdateBalance(wallet.ID.String(), debitWallet)

	if err != nil {
		return err
	}

	card, err := ts.CardRepository.GetBy(body.Data.Object.Card.Id)

	if err != nil {
		return err
	}

	feeTransaction := domain.Transaction{
		Company:           wallet.Company,
		Wallet:            wallet.ID,
		Card:              card.ID,
		PartnerCardID:     card.PartnerCardID,
		Customer:          customer.ID,
		PartnerCustomerID: customer.PartnerCustomerID,
		Debit:             fee,
		Note:              fmt.Sprintf("%v was debitted for transaction fee", fee),
		ReferenceID:       body.Data.Id,
		Status:            domain.PendingStatus,
		Entry:             domain.DebitEntry,
		Channel:           domain.TransactionChannel(payload.TransactionMetadata.Channel),
		Type:              domain.FeeType,
		CardType:          domain.CardType(payload.Card.Type),
	}

	transaction := domain.Transaction{
		Company:           wallet.Company,
		Wallet:            wallet.ID,
		Card:              card.ID,
		PartnerCardID:     card.PartnerCardID,
		Customer:          customer.ID,
		PartnerCustomerID: customer.PartnerCustomerID,
		Debit:             fee,
		Note:              fmt.Sprintf("%v was debitted for transaction", payload.PendingRequest.Amount),
		ReferenceID:       payload.Id,
		Status:            domain.PendingStatus,
		Entry:             domain.DebitEntry,
		Channel:           domain.TransactionChannel(payload.TransactionMetadata.Channel),
		Type:              domain.WithdrawalType,
		CardType:          domain.CardType(payload.Card.Type),
	}

	err = ts.TransactionRepository.Persist(&feeTransaction)
	if err != nil {
		return err
	}

	err = ts.TransactionRepository.Persist(&transaction)
	if err != nil {
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
