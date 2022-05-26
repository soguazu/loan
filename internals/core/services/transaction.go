package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"errors"
	"fmt"
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
	payload := body.Data.Object
	customerEntity := domain.Customer{
		PartnerCustomerID: payload.Customer.Id,
	}

	customer, err := ts.CustomerRepository.GetBy(customerEntity)
	if err != nil {
		return err
	}

	wallet, err := ts.WalletRepository.GetByCompany(customer.Company.String())

	if err != nil {
		return err
	}

	if strings.ToLower(strings.TrimSpace(body.Type)) == "transaction.created" {
		fmt.Println("Just got here", body, "transaction.created")

		err := ts.ProcessTransactionState(body, wallet)

		if err != nil {
			return err
		}

		return nil
	} else if strings.ToLower(strings.TrimSpace(body.Type)) == "authorization.request" {

		var (
			chargeIdentify common.PricingIdentifier
		)

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

		totalAmountInKoBo := utils.ToMinorUnit(totalAmount)

		entryType := string(domain.DebitEntry)
		
		debitWallet := common.UpdateWalletRequest{
			CreditLimit:     &wallet.CreditLimit,
			PreviousBalance: &wallet.PreviousBalance,
			CurrentSpending: &wallet.CurrentSpending,
			Payment:         &totalAmountInKoBo,
			Entry:           &entryType,
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
			ReferenceID:       body.Id,
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
			Debit:             float64(payload.PendingRequest.Amount),
			Note:              fmt.Sprintf("%v was debitted for transaction", payload.PendingRequest.Amount),
			ReferenceID:       body.Id,
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

	return errors.New("invalid webhook")
}

func (ts *transactionService) ProcessTransactionState(body *common.CreateTransactionRequest, wallet *domain.Wallet) error {
	if strings.ToLower(strings.TrimSpace(body.Data.Object.Status)) == "approved" {
		transactionEntity := domain.Transaction{
			ReferenceID: body.Id,
		}

		transactions, err := ts.TransactionRepository.GetBy(transactionEntity)

		if err != nil {
			return err
		}

		for _, transaction := range transactions {
			transaction.Status = domain.SuccessStatus
			ts.TransactionRepository.Persist(&transaction)
		}

		return nil

	} else if strings.ToLower(strings.TrimSpace(body.Data.Object.Status)) == "failed" {
		var charges float64
		var err error

		transactionEntity := domain.Transaction{
			ReferenceID: body.Id,
		}

		transactions, err := ts.TransactionRepository.GetBy(transactionEntity)

		if err != nil {
			return err
		}

		for _, transaction := range transactions {
			transaction.Status = domain.FailedStatus
			charges += transaction.Debit
			ts.TransactionRepository.Persist(&transaction)

			newTransaction := &domain.Transaction{
				Company:           transaction.Company,
				Wallet:            wallet.ID,
				Card:              transaction.Card,
				PartnerCardID:     transaction.PartnerCardID,
				Customer:          transaction.Customer,
				PartnerCustomerID: transaction.PartnerCustomerID,
				Debit:             transaction.Debit,
				Note:              fmt.Sprintf("%v was refunded for failed transaction", transaction.Debit),
				ReferenceID:       body.Id,
				Status:            domain.SuccessStatus,
				Entry:             domain.CreditEntry,
				Channel:           transaction.Channel,
				Type:              domain.RefundType,
				ParentID:          transaction.ID.String(),
			}

			ts.TransactionRepository.Persist(newTransaction)

		}

		chargesInKobo := utils.ToMinorUnit(charges)

		ts.WalletService.CreditWallet(wallet, chargesInKobo)
	}

	return errors.New("invalid webhook")
}

func (ts *transactionService) UpdateTransaction(id string, body common.UpdateTransactionRequest) (*domain.Transaction, error) {
	transaction, err := ts.TransactionRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if body.Receipt != nil {
		transaction.Receipt = *body.Receipt
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
