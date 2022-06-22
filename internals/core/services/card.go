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
	"net/http"
	"strings"
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
	PANRepository         ports.IPANRepository
	FeeRepository         ports.IFeeRepository
	logger                *log.Logger
}

// NewCardService function create a new instance for service
func NewCardService(cr ports.ICardRepository, csr ports.ICustomerRepository,
	ar ports.IAddressRepository, cmr ports.ICompanyRepository, fr ports.IFeeRepository,
	ws ports.IWalletService, tr ports.ITransactionRepository, pr ports.IPANRepository,
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
		PANRepository:         pr,
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
	var chargesInKobo *int64

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

	if len(address) < 1 {
		return nil, errors.New("company address not set")
	}

	wallet, err := cs.WalletRepository.GetBy(walletEntity)

	if len(wallet) < 1 {
		return nil, errors.New("account has not been integrated")
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
		Individual: common.Individual{
			FirstName: body.User.FirstName,
			LastName:  body.User.LastName,
		},
		Status: body.Status,
		Type:   common.CustomerType,
		Name:   fmt.Sprintf("%v %v", body.User.FirstName, body.User.LastName),
		Phone:  body.User.Phone,
		Email:  body.User.Email,
		BillingAddress: common.BillingAddress{
			Line1:      address[0].Address,
			City:       address[0].City,
			State:      address[0].State,
			PostalCode: address[0].PostalCode,
			Country:    address[0].Country,
		},
	}

	if strings.ToLower(body.Type) == "virtual" {
		chargesIdentifier = common.VirtualCardIdentifier
		chargesInKobo, err = cs.GetAllCharges(chargesIdentifier, &wallet[0])
	} else if strings.ToLower(body.Type) == "physical" {
		chargesIdentifier = common.PhysicalCardIdentifier
		chargesInKobo, err = cs.GetAllCharges(chargesIdentifier, &wallet[0])
	}

	if err != nil {
		return nil, err
	}

	isBalanceSufficient := cs.BalanceSufficiency(&wallet[0], *chargesInKobo)

	if !isBalanceSufficient {
		return nil, errors.New("insufficient available credit")
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

	cardEntity, err := cs.ReturnCardObj(&body, sudoCustomer)

	if err != nil {
		return nil, err
	}

	fmt.Println(cardEntity)

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

	fmt.Println(response)

	if response.StatusCode != 200 {
		return nil, errors.New("error occurred creating card partner")
	}

	_, err = cs.WalletService.DebitWallet(&wallet[0], *chargesInKobo)

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
		Partner:           common.Partner,
		Summary:           body.Summary,
		Customer:          customer.ID,
		SpendingControls: domain.SpendingControls{
			Channels: domain.Channels{
				Pos:    response.Data.SpendingControls.Channels.Pos,
				Web:    response.Data.SpendingControls.Channels.Web,
				Atm:    response.Data.SpendingControls.Channels.Atm,
				Mobile: response.Data.SpendingControls.Channels.Mobile,
			},
			SpendingLimits: domain.SpendingLimits{
				Amount:   response.Data.SpendingControls.SpendingLimits[0].Amount,
				Interval: response.Data.SpendingControls.SpendingLimits[0].Interval,
			},
		},
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

	byteResponse, err := CardRequest.CHANGE("POST", "customers", byteBody)

	var response common.CreateSudoCustomerResponse

	err = json.Unmarshal(byteResponse, &response)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("error occurred creating customer partner")
	}

	return &response, nil
}

func (cs *cardService) ReturnCardObj(body *common.CreateCardRequest,
	sudoCustomer *common.CreateSudoCustomerResponse) (*common.CreateSudoCardRequest, error) {
	var card common.CreateSudoCardRequest

	if strings.ToLower(body.Type) == "virtual" {
		card = common.CreateSudoCardRequest{
			Type:             body.Type,
			Brand:            cs.Capitalize(strings.ToLower(body.Brand)),
			Currency:         common.Currency,
			Status:           body.Status,
			CustomerID:       sudoCustomer.Data.ID,
			FundingSourceID:  config.Instance.FundingSource,
			SpendingControls: body.SpendingControls,
		}
	} else if strings.ToLower(body.Type) == "physical" {
		var pan *domain.PAN
		var sudoPAN *common.SudoPANNumber
		var err error
		var number string

		if config.Instance.Env == "development" {
			sudoPAN, err = cs.GetPANSudo()

			if sudoPAN == nil {
				return nil, errors.New("no PAN available on partner")
			}

			if err != nil {
				return nil, err
			}

			fmt.Println(sudoPAN, "$$$$$")
			number = sudoPAN.Data.Number

		} else if config.Instance.Env == "production" {
			pan, err = cs.PANRepository.GetFirstOne()

			if pan == nil {
				return nil, errors.New("no PAN available")
			}

			if err != nil {
				return nil, err
			}

			number = pan.Number
		}

		card = common.CreateSudoCardRequest{
			Type:             body.Type,
			Brand:            cs.Capitalize(strings.ToLower(body.Brand)),
			Currency:         common.Currency,
			Status:           body.Status,
			CustomerID:       sudoCustomer.Data.ID,
			FundingSourceID:  config.Instance.FundingSource,
			Number:           number,
			SpendingControls: body.SpendingControls,
		}

	} else {
		return nil, errors.New("invalid card type")
	}

	return &card, nil
}

func (cs *cardService) Capitalize(value string) string {
	return strings.ToUpper(string(value[0])) + value[1:]
}

func (cs *cardService) UpdateCard(id string, body common.UpdateSudoCardRequest) (*domain.Card, error) {
	Headers := map[string]string{
		"Accept":        "application/json; charset=utf-8",
		"Authorization": "Bearer " + config.Instance.SudoAPIKey,
		"Content-Type":  "application/json",
	}

	CardRequest = utils.Client{
		BaseURL: config.Instance.SudoBaseURL,
		Header:  Headers,
	}

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

	body.SpendingControls.AllowedCategories = []string{}
	body.SpendingControls.BlockedCategories = []string{}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	fmt.Println("got here", string(byteBody))

	url := fmt.Sprintf("%v/%v", "cards", card.PartnerCardID)

	fmt.Println(url, "url")

	byteResponse, err := CardRequest.CHANGE(http.MethodPut, url, byteBody)

	if err != nil {
		fmt.Println("E HAPPEN")
		return nil, err
	}

	var response common.ProcessCardUpdate

	err = json.Unmarshal(byteResponse, &response)

	if response.StatusCode != 200 {
		return nil, errors.New("error occurred updating card partner")
	}

	err = cs.CardRepository.Persist(card)

	if err != nil {
		fmt.Println("E HAPPEN 2")

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
	Headers := map[string]string{
		"Accept":        "application/json; charset=utf-8",
		"Authorization": "Bearer " + config.Instance.SudoAPIKey,
		"Content-Type":  "application/json",
	}

	CardRequest = utils.Client{
		BaseURL: config.Instance.SudoBaseURL,
		Header:  Headers,
	}

	card, err := cs.CardRepository.GetByID(id)

	if err != nil {
		return err
	}

	isValid := body.Status != "active" && body.Status != "inactive" && body.Status != "cancelled"

	if isValid {
		return errors.New("invalid status")
	}

	var listOfSpendingLimits = []common.SpendingLimits{
		common.SpendingLimits{
			Amount:   card.SpendingControls.SpendingLimits.Amount,
			Interval: card.SpendingControls.SpendingLimits.Interval,
		},
	}

	request := common.CancelCardRequest{
		Status: body.Status,
		SpendingControls: common.SpendingControls{
			SpendingLimits: listOfSpendingLimits,
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

	byteResponse, err := CardRequest.CHANGE(http.MethodPut, url, byteBody)

	if err != nil {
		return err
	}

	var response common.ProcessCardUpdate

	err = json.Unmarshal(byteResponse, &response)

	if response.StatusCode != 200 {
		return errors.New("error occurred updating card partner")
	}

	card.Status = body.Status

	err = cs.CardRepository.Delete(card.ID.String())

	if err != nil {
		return err
	}

	return nil

}

func (cs *cardService) ChangeCardPin(id string, body common.ChangeCardPinRequest) error {

	Headers := map[string]string{
		"Accept":        "application/json; charset=utf-8",
		"Authorization": "Bearer " + config.Instance.SudoAPIKey,
		"Content-Type":  "application/json",
	}

	CardRequest = utils.Client{
		BaseURL: config.Instance.SudoBaseURL,
		Header:  Headers,
	}

	card, err := cs.CardRepository.GetByID(id)
	if err != nil {
		return err
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v/%v/%v", "cards", card.PartnerCardID, "pin")

	byteResponse, err := CardRequest.CHANGE("PUT", url, byteBody)

	if err != nil {
		return err
	}

	var response common.ProcessCardUpdate

	err = json.Unmarshal(byteResponse, &response)

	fmt.Println(response)

	if response.StatusCode != 200 {
		return errors.New("error occurred updating card partner")
	}

	return nil
}

func (cs *cardService) BalanceSufficiency(wallet *domain.Wallet, chargesInKobo int64) bool {
	if wallet.AvailableCredit > chargesInKobo {
		return true
	}
	return false
}

func (cs *cardService) GetAllCharges(identifiers []common.PricingIdentifier, wallet *domain.Wallet) (*int64, error) {
	var charges float64
	for _, identifier := range identifiers {
		fee, err := cs.FeeRepository.GetByIdentifier(string(identifier))
		if err != nil {
			return nil, err
		}
		charges += fee.Fee
	}

	chargesInKobo := utils.ToMinorUnit(charges)

	return &chargesInKobo, nil
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

func (cs *cardService) AddPAN(body common.AddPANRequest) error {
	var pans []domain.PAN

	for _, pan := range body.Numbers {
		pans = append(pans, domain.PAN{
			Number: pan,
		})
	}

	err := cs.PANRepository.BatchInsert(pans)
	if err != nil {
		cs.logger.Error(err)
		return err
	}
	return nil
}

func (cs *cardService) GetSinglePAN() (*domain.PAN, error) {
	pan, err := cs.PANRepository.GetFirstOne()
	if err != nil {
		return nil, err
	}
	return pan, nil
}

func (cs *cardService) DeletePAN(id string) error {
	err := cs.PANRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *cardService) GetPANSudo() (*common.SudoPANNumber, error) {
	byteResponse, err := CardRequest.GET("GET", "cards/simulator/generate", nil)

	if err != nil {
		return nil, err
	}

	var response common.SudoPANNumber

	err = json.Unmarshal(byteResponse, &response)

	fmt.Println(response)

	if response.StatusCode != 200 {
		return nil, errors.New("error occurred updating card partner")
	}
	return &response, nil
}
