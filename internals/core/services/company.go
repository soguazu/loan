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
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// Request to create API request
var Request utils.Client

type companyService struct {
	CompanyRepository             ports.ICompanyRepository
	CompanyProfileRepository      ports.ICompanyProfileRepository
	WalletRepository              ports.IWalletRepository
	CreditLimitIncreaseRepository ports.ICreditLimitRequestRepository
	logger                        *log.Logger
}

// NewCompanyService function create a new instance for service
func NewCompanyService(cr ports.ICompanyRepository,
	cpr ports.ICompanyProfileRepository,
	wr ports.IWalletRepository,
	cli ports.ICreditLimitRequestRepository,
	l *log.Logger) ports.ICompanyService {
	return &companyService{
		CompanyRepository:             cr,
		CompanyProfileRepository:      cpr,
		WalletRepository:              wr,
		CreditLimitIncreaseRepository: cli,
		logger:                        l,
	}
}

func (c *companyService) GetCompanyByID(id string) (*domain.Company, error) {
	company, err := c.CompanyRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c *companyService) GetCompany(filter interface{}) ([]domain.Company, error) {
	companies, err := c.CompanyRepository.GetBy(filter)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (c *companyService) GetAllCompany(pagination *utils.Pagination) (*utils.Pagination, error) {
	companies, err := c.CompanyRepository.Get(pagination)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (c *companyService) CreateCompany(company *domain.Company) error {
	var entity []domain.Company
	entity, err := c.CompanyRepository.GetBy(domain.Company{Owner: company.Owner, Name: company.Name})
	if err != nil {
		return err
	}

	if len(entity) > 0 {
		return errors.New("already exist")
	}

	err = c.CompanyRepository.Persist(company)

	if err != nil {
		c.logger.Error(err)
		return err
	}

	return nil
}

func (c *companyService) DeleteCompany(id string) error {
	err := c.CompanyRepository.Delete(id)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	return nil
}

func (c *companyService) UpdateCompany(params common.GetByIDRequest, body common.UpdateCompanyRequest) (*domain.Company, error) {
	company, err := c.CompanyRepository.GetByID(params.ID)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	if body.Type != nil {
		company.Type = *body.Type
	}

	if body.Website != nil {
		company.Website = *body.Website
	}

	if body.Name != nil {
		company.Name = *body.Name
	}

	if body.NoOfEmployee != nil {
		company.NoOfEmployee = *body.NoOfEmployee
	}

	if body.FundingSource != nil {
		company.FundingSource = *body.FundingSource
	}

	err = c.CompanyRepository.Persist(company)

	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return company, nil
}

func (c *companyService) UnderWriting(id string) (*common.UnderWritingResponse, error) {
	Headers := map[string]string{
		"Accept":        "application/json; charset=utf-8",
		"Authorization": "Bearer " + config.Instance.OkraSecret,
		"Content-Type":  "application/json",
	}

	Request = utils.Client{
		BaseURL: config.Instance.OkraBaseURL,
		Header:  Headers,
	}

	companyUUID, err := uuid.FromString(id)

	if err != nil {
		return nil, err
	}

	companyProfileEntity := domain.CompanyProfile{
		Company: companyUUID,
	}

	companyProfile, err := c.CompanyProfileRepository.GetBy(companyProfileEntity)

	if err != nil {
		return nil, err
	}

	walletEntity := domain.Wallet{
		Company: companyUUID,
	}

	wallet, err := c.WalletRepository.GetBy(walletEntity)
	if err != nil {
		return nil, err
	}

	err = c.KYCCheck(&companyProfile[0])

	if err != nil {
		return nil, err
	}

	cashBalancePoint, averageCashBalance, err := c.CashBalanceCheck(&wallet[0])

	if err != nil {
		return nil, err
	}

	totalCredit, totalDebit, transactionsPoint, err := c.TransactionsCheck(&wallet[0])

	if err != nil {
		return nil, err
	}

	financialWellnessPoint, err := c.FinancialWellnessCheck(totalCredit, totalDebit, averageCashBalance)

	if err != nil {
		return nil, err
	}

	yearsOfOperationPoint, err := c.YearsOfOperationCheck(&companyProfile[0])

	if err != nil {
		return nil, err
	}

	totalPoint := cashBalancePoint + transactionsPoint + financialWellnessPoint + yearsOfOperationPoint

	var creditLimit float64

	if totalPoint >= 10 {
		creditLimit, err = c.CalculateCreditLimit(totalPoint, averageCashBalance)

		if err != nil {
			return nil, err
		}
	}

	wallet[0].CreditLimit = utils.ToMinorUnit(creditLimit)

	err = c.WalletRepository.Persist(&wallet[0])
	if err != nil {
		return nil, err
	}

	return &common.UnderWritingResponse{
		CreditLimit: creditLimit,
		TotalPoint:  totalPoint,
	}, nil
}

func (c *companyService) KYCCheck(companyProfile *domain.CompanyProfile) error {
	company, err := c.CompanyRepository.GetByID(companyProfile.Company.String())
	if err != nil {
		return err
	}

	body := common.KYCCheckRequest{
		CompanyName: company.Name,
		RCNumber:    companyProfile.RCNumber,
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	byteResponse, err := Request.CHANGE("POST", "products/kyc/rc-verify", byteBody)

	var response common.KYCCheckResponse

	err = json.Unmarshal(byteResponse, &response)

	if err != nil {
		return err
	}

	if strings.ToLower(response.Data.Details.CompanyName) != strings.ToLower(company.Name) {
		return errors.New("company name does not match")
	}

	return nil
}

func (c *companyService) CashBalanceCheck(wallet *domain.Wallet) (int32, float64, error) {
	const MinCashBalance = 1000000
	From := time.Now().AddDate(0, -7, 0).Format("2006-01-02")
	To := time.Now().AddDate(0, -1, 0).Format("2006-01-02")

	body := common.CashBalanceRequest{
		From:      From,
		To:        To,
		AccountID: wallet.AccountID,
		Page:      1,
		Limit:     2000000,
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return 0, 0, err
	}

	byteResponse, err := Request.CHANGE("POST", "transactions/balance/process", byteBody)

	var response common.CashBalanceResponse

	err = json.Unmarshal(byteResponse, &response)

	if err != nil {
		return 0, 0, err
	}

	var totalBalance float64
	for _, transaction := range response.Data.Data[0].Transactions {
		totalBalance += transaction.BankBalance
	}
	fmt.Println(len(response.Data.Data[0].Transactions), "total cashback")

	averageBalance := totalBalance / float64(len(response.Data.Data[0].Transactions))

	if averageBalance < MinCashBalance {
		return 1, averageBalance, nil
	} else if averageBalance == MinCashBalance {
		return 3, averageBalance, nil
	}
	return 5, averageBalance, nil
}

func (c *companyService) TransactionsCheck(wallet *domain.Wallet) (float64, float64, int32, error) {

	From := time.Now().AddDate(0, -7, 0).Format("2006-01-02")
	To := time.Now().AddDate(0, -1, 0).Format("2006-01-02")

	body := common.TransactionsCheckRequest{
		From:       From,
		To:         To,
		CustomerID: wallet.CustomerID,
		Page:       1,
		Limit:      2000000,
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return 0, 0, 0, err
	}

	byteResponse, err := Request.CHANGE("POST", "transactions/getByCustomerDate", byteBody)

	if err != nil {
		return 0, 0, 0, err
	}

	var response common.TransactionCheckResponse

	err = json.Unmarshal(byteResponse, &response)

	if err != nil {
		return 0, 0, 0, err
	}

	var totalCredit, totalDebit float64

	for _, transaction := range response.Data.Transaction {
		if transaction.Credit != 0 {
			totalCredit += transaction.Credit
		} else {
			totalDebit += transaction.Debit
		}
	}

	fmt.Println(len(response.Data.Transaction), "transactions")

	totalPoints, err := c.CalculateEIRation(totalCredit, totalDebit)

	if err != nil {
		return 0, 0, 0, err
	}

	return totalCredit, totalDebit, totalPoints, nil
}

func (c *companyService) FinancialWellnessCheck(totalCredit, totalDebit, averageCashBack float64) (int32, error) {
	// operational cash-flow ratio
	const MinOCR float64 = 1.0
	financialWellness := (totalCredit + averageCashBack) / totalDebit

	if financialWellness < MinOCR {
		return 0, nil
	} else if financialWellness >= MinOCR && financialWellness <= 1.5 {
		return 3, nil
	}
	return 5, nil
}

func (c *companyService) YearsOfOperationCheck(companyProfile *domain.CompanyProfile) (int32, error) {
	yearOfOperations := time.Now().Year() - companyProfile.YearsInOperation

	switch {
	case yearOfOperations > 0 && yearOfOperations <= 2:
		return 1, nil
	case yearOfOperations >= 3 && yearOfOperations <= 5:
		return 3, nil
	case yearOfOperations > 6:
		return 5, nil
	default:
		return 0, errors.New("invalid years of operations")
	}
}

func (c *companyService) CalculateCreditLimit(totalPoint int32, averageCashBalance float64) (float64, error) {
	if totalPoint < 1 || averageCashBalance < 1 {
		return 0, errors.New("invalid total point or average cash balance")
	}

	var creditLimitPercent float64 = 70 / 100
	creditLimit := creditLimitPercent * averageCashBalance
	fmt.Println(creditLimit, "CREDIT LIMIT")
	return creditLimit, nil
}

func (c *companyService) CalculateEIRation(totalCredit, totalDebit float64) (int32, error) {
	if totalCredit < 1 || totalDebit < 1 {
		return 0, errors.New("invalid total credit or total debit")
	}

	eiRatio := totalCredit / totalDebit * 100
	if eiRatio >= 60 {
		return 0, nil
	} else if eiRatio >= 50 && eiRatio < 59 {
		return 3, nil
	}
	return 5, nil

}

func (c *companyService) RequestCreditLimitIncrease(id string, body *domain.CreditIncrease) error {
	ID, err := uuid.FromString(id)

	if err != nil {
		return err
	}

	walletEntity := domain.Wallet{
		Company: ID,
	}

	wallet, err := c.WalletRepository.GetBy(walletEntity)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	company, err := c.CompanyRepository.GetByID(id)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	body.Wallet = wallet[0].ID
	body.Company = company.ID
	body.Owner = company.Owner

	err = c.CreditLimitIncreaseRepository.Persist(body)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyService) UpdateRequestCreditLimitIncrease(params common.GetByIDRequest, body common.UpdateCreditLimitIncreaseRequest) error {
	ID, err := uuid.FromString(params.ID)

	if err != nil {
		return err
	}

	creditLimitRequest := domain.CreditIncrease{
		Company: ID,
	}

	creditLimit, err := c.CreditLimitIncreaseRepository.GetBy(creditLimitRequest)

	if err != nil {
		c.logger.Error(err)
		return err
	}

	if body.DesiredCreditLimit != nil {
		creditLimit.DesiredCreditLimit = *body.DesiredCreditLimit
	}

	if body.Reason != nil {
		creditLimit.Reason = *body.Reason
	}

	if body.Status != nil {
		creditLimit.Status = *body.Status
	}

	err = c.CreditLimitIncreaseRepository.Persist(creditLimit)

	if err != nil {
		c.logger.Error(err)
		return err
	}
	return nil
}
