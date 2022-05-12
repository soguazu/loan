package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/config"
	"core_business/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// CardRequest to create API request
var CardRequest utils.Client

type cardService struct {
	CardRepository        ports.ICardRepository
	CompanyRepository     ports.ICompanyRepository
	CustomerRepository    ports.ICustomerRepository
	WalletRepository      ports.IWalletRepository
	WalletService         ports.IWalletService
	AddressRepository     ports.IAddressRepository
	TransactionRepository ports.ITransactionRepository
	FeeRepository         ports.IFeeRepository
	logger                *log.Logger
}

// NewCardService function create a new instance for service
func NewCardService(cr ports.ICardRepository, csr ports.ICustomerRepository,
	ar ports.IAddressRepository, cmr ports.ICompanyRepository, fr ports.IFeeRepository,
	ws ports.IWalletService, tr ports.ITransactionRepository,
	wr ports.IWalletRepository, l *log.Logger) ports.ICardService {
	return &cardService{
		CardRepository:        cr,
		CompanyRepository:     cmr,
		CustomerRepository:    csr,
		AddressRepository:     ar,
		FeeRepository:         fr,
		WalletService:         ws,
		WalletRepository:      wr,
		TransactionRepository: tr,
		logger:                l,
	}
}

func (cs *cardService) GetCardByID(id string) (*domain.Card, error) {
	card, err := cs.CardRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (cs *cardService) GetCardByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error) {
	cards, err := cs.CardRepository.GetCardByCompanyID(id, pagination)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (cs *cardService) GetAllCard(pagination *utils.Pagination) (*utils.Pagination, error) {
	cards, err := cs.CardRepository.Get(pagination)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (cs *cardService) CreateCard(body common.CreateCardRequest) (*domain.Card, error) {
	var err error
	var chargesIdentifier []common.PricingIdentifier

	Headers := map[string]string{
		"Accept":        "application/json; charset=utf-8",
		"Authorization": "Bearer " + config.Instance.SudoAPIKey,
		"Content-Type":  "application/json",
	}

	CardRequest = utils.Client{
		BaseURL: config.Instance.SudoBaseURL,
		Header:  Headers,
	}

	addressEntity := domain.Address{Company: body.Company}
	walletEntity := domain.Wallet{Company: body.Company}

	address, err := cs.AddressRepository.GetBy(addressEntity)
	if err != nil {
		return nil, err
	}

	wallet, err := cs.WalletRepository.GetBy(walletEntity)

	if err != nil {
		return nil, err
	}

	if body.Type == "virtual" {
		chargesIdentifier = common.VirtualCardIdentifier
		err = cs.GetAllCharges(chargesIdentifier, &wallet[0])
	} else {
		chargesIdentifier = common.PhysicalCardIdentifier
		err = cs.GetAllCharges(chargesIdentifier, &wallet[0])
	}

	if err != nil {
		return nil, err
	}

	company, err := cs.CompanyRepository.GetByID(body.Company.String())
	if err != nil {
		return nil, err
	}

	customerDTO := common.CreateCustomerRequest{
		Company: common.Company{
			Name: company.Name,
		},
		Status: common.CustomerStatus,
		Type:   common.CustomerType,
		Name:   fmt.Sprintf("%v %v", body.FirstName, body.LastName),
		Phone:  body.Phone,
		Email:  body.Email,
		BillingAddress: common.BillingAddress{
			Line1:      address[0].Address,
			City:       address[0].City,
			State:      address[0].State,
			PostalCode: address[0].PostalCode,
			Country:    address[0].Country,
		},
	}

	sudoCustomer, err := cs.CreateSudoCustomer(customerDTO)

	if err != nil {
		return nil, err
	}

	customer := &domain.Customer{
		Company:           company.ID,
		Wallet:            wallet[0].ID,
		PartnerCustomerID: sudoCustomer.Data.ID,
		Address:           sudoCustomer.Data.BillingAddress.Line1,
		City:              sudoCustomer.Data.BillingAddress.City,
		State:             sudoCustomer.Data.BillingAddress.State,
		Country:           sudoCustomer.Data.BillingAddress.Country,
		PostalCode:        sudoCustomer.Data.BillingAddress.PostalCode,
	}

	cardEntity := cs.ReturnCardObj(&body, sudoCustomer)

	byteBody, err := json.Marshal(cardEntity)

	if err != nil {
		return nil, err
	}

	byteResponse, err := CardRequest.CHANGE("POST", "cards", byteBody)

	var response common.CreateSudoCardResponse

	err = json.Unmarshal(byteResponse, &response)

	if err != nil {
		return nil, err
	}

	err = cs.CustomerRepository.Persist(customer)

	if err != nil {
		return nil, err
	}

	card := &domain.Card{
		Company:           company.ID,
		Wallet:            wallet[0].ID,
		PartnerCustomerID: sudoCustomer.Data.ID,
		PartnerCardID:     response.Data.ID,
		Type:              response.Data.Type,
		Brand:             response.Data.Brand,
		Currency:          response.Data.Currency,
		Status:            response.Data.Status,
		Lock:              body.Lock,
		Partner:           common.Partner,
		CardAuth:          body.CardAuth,
		Summary:           body.Summary,
		Customer:          customer.ID,
		//SpendingControls:  cardEntity.SpendingControls,
		Business:    response.Data.Business,
		Account:     response.Data.Account,
		MaskedPan:   response.Data.MaskedPan,
		ExpiryMonth: response.Data.ExpiryMonth,
		ExpiryYear:  response.Data.ExpiryYear,
	}

	err = cs.CardRepository.Persist(card)

	if err != nil {
		return nil, err
	}

	err = cs.LogTransactions(chargesIdentifier, card)

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (cs *cardService) CreateSudoCustomer(customerDTO common.CreateCustomerRequest) (*common.CreateSudoCustomerResponse, error) {
	byteBody, err := json.Marshal(customerDTO)

	if err != nil {
		return nil, err
	}

	byteResponse, err := CardRequest.CHANGE("POST", "customer", byteBody)

	var response common.CreateSudoCustomerResponse

	err = json.Unmarshal(byteResponse, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cs *cardService) ReturnCardObj(body *common.CreateCardRequest,
	sudoCustomer *common.CreateSudoCustomerResponse) common.CreateSudoCardRequest {
	var card common.CreateSudoCardRequest

	if body.Type == "virtual" {
		card = common.CreateSudoCardRequest{
			Type:             body.Type,
			Name:             body.Name,
			Brand:            body.Brand,
			Currency:         body.Currency,
			Status:           body.Status,
			CustomerID:       sudoCustomer.Data.ID,
			FundingSourceID:  config.Instance.FundingSource,
			SpendingControls: body.SpendingControls,
		}
	} else {
		card = common.CreateSudoCardRequest{
			Type:            body.Type,
			Brand:           body.Brand,
			Name:            body.Name,
			Currency:        body.Currency,
			Status:          body.Status,
			CustomerID:      sudoCustomer.Data.ID,
			FundingSourceID: config.Instance.FundingSource,
			Number:          body.Number,
		}
	}

	return card
}

func (cs *cardService) UpdateCard(id string, body common.UpdateSudoCardRequest) (*domain.Card, error) {
	card, err := cs.CardRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if body.Status != nil {
		card.Status = *body.Status
	}

	if body.SpendingControls.SpendingLimits.Amount != nil {
		card.SpendingControls.SpendingLimits.Amount = *body.SpendingControls.SpendingLimits.Amount
	}

	if body.SpendingControls.SpendingLimits.Interval != nil {
		card.SpendingControls.SpendingLimits.Interval = *body.SpendingControls.SpendingLimits.Interval
	}

	if body.SpendingControls.Channels.Atm != nil {
		card.SpendingControls.Channels.Atm = *body.SpendingControls.Channels.Atm
	}

	if body.SpendingControls.Channels.Pos != nil {
		card.SpendingControls.Channels.Pos = *body.SpendingControls.Channels.Pos
	}

	if body.SpendingControls.Channels.Web != nil {
		card.SpendingControls.Channels.Web = *body.SpendingControls.Channels.Web
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%v/%v", "cards", card.PartnerCardID)

	_, err = CardRequest.CHANGE("PUT", url, byteBody)

	if err != nil {
		return nil, err
	}

	err = cs.CardRepository.Persist(card)

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (cs *cardService) LockCard(id string, body common.ActionOnCardRequest) (*domain.Card, error) {
	card, err := cs.CardRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	card.Lock = body.Lock

	err = cs.CardRepository.Persist(card)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (cs *cardService) CancelCard(id string, body common.ChangeCardStatusRequest) error {
	card, err := cs.CardRepository.GetByID(id)

	if err != nil {
		return err
	}

	isValid := body.Status != "active" && body.Status != "inactive" && body.Status != "canceled"

	if isValid {
		return errors.New("invalid status")
	}

	request := common.CancelCardRequest{
		Status: body.Status,
		SpendingControls: common.SpendingControls{
			SpendingLimits: common.SpendingLimits{
				Amount:   card.SpendingControls.SpendingLimits.Amount,
				Interval: card.SpendingControls.SpendingLimits.Interval,
			},
			Channels: common.Channels{
				Pos:    card.SpendingControls.Channels.Pos,
				Web:    card.SpendingControls.Channels.Web,
				Atm:    card.SpendingControls.Channels.Atm,
				Mobile: false,
			},
			AllowedCategories: []string{},
			BlockedCategories: []string{},
		},
	}

	byteBody, err := json.Marshal(request)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v/%v", "cards", card.PartnerCardID)

	_, err = CardRequest.CHANGE("PUT", url, byteBody)

	if err != nil {
		return err
	}

	card.Status = body.Status

	err = cs.CardRepository.Delete(card.ID.String())

	if err != nil {
		return err
	}

	return nil

}

func (cs *cardService) ChangeCardPin(id string, body common.ChangeCardPinRequest) error {
	_, err := cs.CardRepository.GetByID(id)
	if err != nil {
		return err
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v/%v/%v", "cards", body.OldPin, id)

	_, err = CardRequest.CHANGE("PUT", url, byteBody)

	if err != nil {
		return err
	}

	return nil
}

func (cs *cardService) GetAllCharges(identifiers []common.PricingIdentifier, wallet *domain.Wallet) error {
	var charges float64
	for _, identifier := range identifiers {
		fee, err := cs.FeeRepository.GetByIdentifier(string(identifier))
		if err != nil {
			return err
		}
		charges += fee.Fee
	}

	chargesInKobo := utils.ToMinorUnit(charges)

	if wallet.AvailableCredit < chargesInKobo {
		return errors.New("insufficient available credit")
	}

	_, err := cs.DebitWallet(wallet, charges)
	if err != nil {
		return err
	}

	return nil
}

func (cs *cardService) LogTransactions(identifiers []common.PricingIdentifier, card *domain.Card) error {
	var chargedFor string
	var transactionType domain.TransactionType

	for _, identifier := range identifiers {
		fee, err := cs.FeeRepository.GetByIdentifier(string(identifier))
		if err != nil {
			return err
		}

		if identifier == common.CardCreation {
			chargedFor = "purchase"
			transactionType = domain.CardCreationType
		} else {
			chargedFor = "shipping"
			transactionType = domain.ShippingType
		}
		transaction := &domain.Transaction{
			Company:           card.Company,
			Wallet:            card.Wallet,
			Card:              card.ID,
			Customer:          card.Customer,
			PartnerCustomerID: card.PartnerCustomerID,
			Debit:             fee.Fee,
			Note:              fmt.Sprintf("debited for card %v", chargedFor),
			Status:            domain.SuccessStatus,
			Entry:             domain.DebitEntry,
			Channel:           domain.WebChannel,
			Type:              transactionType,
		}

		err = cs.TransactionRepository.Persist(transaction)

		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *cardService) DebitWallet(wallet *domain.Wallet, charges float64) (*domain.Wallet, error) {
	transactionType := string(common.DebitTransaction)
	walletEntity := common.UpdateWalletRequest{
		CreditLimit:     &wallet.CreditLimit,
		PreviousBalance: &wallet.PreviousBalance,
		CurrentSpending: &wallet.CurrentSpending,
		Type:            &transactionType,
		Payment:         &charges,
	}

	wallet, err := cs.WalletService.UpdateBalance(wallet.ID.String(), walletEntity)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
