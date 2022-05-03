package handlers

import (
	"bytes"
	"core_business/internals/common"
	"core_business/internals/core/services"
	"core_business/internals/repositories"
	datastore "core_business/pkg/database"
	"core_business/pkg/logger"
	"core_business/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
)

var (
	db                  = datastore.NewSqliteDatabase()
	DBConnection        = db.ConnectDB(filepath.Join("..", "..", "evea.db"))
	logging             = logger.NewLogger(log.New()).MakeLogger(filepath.Join("..", "..", "logs", "info"), true)
	companyRepository   = repositories.NewCompanyRepository(DBConnection)
	walletRepository    = repositories.NewWalletRepository(DBConnection)
	creditLimitIncrease = repositories.NewCreditLimitRequestRepository(DBConnection)
	companyService      = services.NewCompanyService(companyRepository, companyProfileRepository, walletRepository, creditLimitIncrease, logging)
	handler             = NewCompanyHandler(companyService, logging, "Company")
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCompanyHandler_GetAllCompany(t *testing.T) {
	r := SetupRouter()
	r.GET("/v1/company", handler.GetAllCompany)

	request, err := http.NewRequest("GET", "/v1/company", nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var company common.GetCompanyResponse

	err = json.Unmarshal([]byte(response.Body.String()), &company)
	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
}

func createCompany(t *testing.T) *common.CreateCompanyResponse {
	r := SetupRouter()
	r.POST("/v1/company", handler.CreateCompany)
	entity := common.CreateCompanyRequest{
		Company:       (&utils.Faker{}).RandomUUID(),
		Owner:         (&utils.Faker{}).RandomObjectID(),
		Name:          (&utils.Faker{}).RandomName(),
		Website:       (&utils.Faker{}).RandomWebsite(),
		Type:          (&utils.Faker{}).RandomType(),
		FundingSource: (&utils.Faker{}).RandomFundSource(),
		NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
	}

	jsonValue, _ := json.Marshal(entity)
	request, err := http.NewRequest("POST", "/v1/company", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error occurred decoding")
		return nil
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var company *common.CreateCompanyResponse

	_ = json.Unmarshal([]byte(response.Body.String()), &company)

	require.Equal(t, http.StatusCreated, response.Code)

	return company
}

func TestCompanyHandler_CreateCompany(t *testing.T) {
	company := createCompany(t)
	require.NotEmpty(t, company)

	fmt.Println(company.Data.ID)
}

func TestCompanyHandler_GetCompanyByID(t *testing.T) {
	company := createCompany(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", company.Data.ID.String())

	r.GET("/v1/company/:id", handler.GetCompanyByID)

	request, err := http.NewRequest("GET", "/v1/company/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.GetCompanyResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, company.Data.ID, resp.ID)
	require.Equal(t, company.Data.Owner, resp.Owner)
	require.Equal(t, company.Data.Name, resp.Name)
	require.Equal(t, company.Data.Type, resp.Type)
	require.Equal(t, company.Data.FundingSource, resp.FundingSource)
	require.Equal(t, company.Data.NoOfEmployee, resp.NoOfEmployee)
}

func TestCompanyHandler_DeleteCompany(t *testing.T) {
	r := SetupRouter()
	company := createCompany(t)

	r.DELETE("/v1/company/:id", handler.DeleteCompany)

	endpoint := fmt.Sprintf("/v1/company/%v", company.Data.ID.String())

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestCompanyHandler_UpdateCompany(t *testing.T) {
	r := SetupRouter()
	company := createCompany(t)

	r.PATCH("/v1/company/:id", handler.UpdateCompany)

	endpoint := fmt.Sprintf("/v1/company/%v", company.Data.ID.String())

	companyName := "CITY BANK"
	companyType := "Jay"
	body := common.UpdateCompanyRequest{
		Name: &companyName,
		Type: &companyType,
	}

	bytePayload, err := json.Marshal(body)
	if err != nil {
		return
	}

	request, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(bytePayload))

	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusOK, response.Code)
}
