package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"core_business/internals/common"
	"core_business/internals/core/services"
	"core_business/internals/repositories"
	"core_business/pkg/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	businessPartnerRepository = repositories.NewBusinessPartnerRepository(DBConnection)
	businessPartnerService    = services.NewBusinessPartnerService(businessPartnerRepository, logging)
	busPartnerHandler         = NewBusinessPartnerHandler(businessPartnerService, logging, "Business partner")
)

func TestBusinessPartnerHandler_GetAllBusinessPartner(t *testing.T) {
	r := SetupRouter()
	r.GET("/v1/business_partner", busPartnerHandler.GetAllBusinessPartner)

	request, err := http.NewRequest("GET", "/v1/business_partner", nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var businessPartner common.BusinessPartnerDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &businessPartner)
	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
}

func createBusinessPartner(t *testing.T) *common.BusinessPartnerDataResponse {
	r := SetupRouter()
	r.POST("/v1/business_partner", busPartnerHandler.CreateBusinessPartner)
	entity := common.CreateBusinessPartnerRequest{
		Company: (&utils.Faker{}).RandomUUID(),
		Name:    (&utils.Faker{}).RandomName(),
		Phone:   (&utils.Faker{}).RandomString(11),
	}

	jsonValue, _ := json.Marshal(entity)
	request, err := http.NewRequest("POST", "/v1/business_partner", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error occurred decoding")
		return nil
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var businessPartner *common.BusinessPartnerDataResponse

	_ = json.Unmarshal([]byte(response.Body.String()), &businessPartner)

	require.Equal(t, http.StatusCreated, response.Code)

	return businessPartner
}

func TestBusinessPartnerHandler_CreateBusinessPartner(t *testing.T) {
	businessPartner := createBusinessPartner(t)
	require.NotEmpty(t, businessPartner)
}

func TestBusinessPartnerHandler_GetBusinessPartnerByID(t *testing.T) {
	businessPartner := createBusinessPartner(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", businessPartner.Data.ID.String())

	r.GET("/v1/business_partner/:id", busPartnerHandler.GetBusinessPartnerByID)

	request, err := http.NewRequest("GET", "/v1/business_partner/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.BusinessPartnerDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, businessPartner.Data.ID, resp.Data.ID)
	require.Equal(t, businessPartner.Data.Name, resp.Data.Name)
	require.Equal(t, businessPartner.Data.Company, resp.Data.Company)
	require.Equal(t, businessPartner.Data.Phone, resp.Data.Phone)
}

func TestBusinessPartnerHandler_GetBusinessPartnerByCompanyID(t *testing.T) {
	businessPartner := createBusinessPartner(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", businessPartner.Data.Company.String())

	r.GET("/v1/business_partner/company/:id", busPartnerHandler.GetBusinessPartnerByCompanyID)

	request, err := http.NewRequest("GET", "/v1/business_partner/company/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.BusinessPartnerDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, businessPartner.Data.ID, resp.Data.ID)
	require.Equal(t, businessPartner.Data.Name, resp.Data.Name)
	require.Equal(t, businessPartner.Data.Company, resp.Data.Company)
	require.Equal(t, businessPartner.Data.Phone, resp.Data.Phone)
}

func TestBusinessPartnerHandler_DeleteBusinessPartner(t *testing.T) {
	r := SetupRouter()
	businessPartner := createBusinessPartner(t)

	r.DELETE("/v1/business_partner/:id", busPartnerHandler.DeleteBusinessPartner)

	endpoint := fmt.Sprintf("/v1/business_partner/%v", businessPartner.Data.ID.String())

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestBusinessPartnerHandler_UpdateBusinessPartner(t *testing.T) {
	r := SetupRouter()
	businessPartner := createBusinessPartner(t)

	r.PATCH("/v1/business_partner/:id", busPartnerHandler.UpdateBusinessPartner)

	endpoint := fmt.Sprintf("/v1/business_partner/%v", businessPartner.Data.ID.String())
	businessPartnerName := (&utils.Faker{}).RandomName()
	businessPartnerPhone := (&utils.Faker{}).RandomString(11)

	body := common.UpdateBusinessPartnerRequest{
		Name:  &businessPartnerName,
		Phone: &businessPartnerPhone,
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
