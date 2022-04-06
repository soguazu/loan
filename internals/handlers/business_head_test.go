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
	businessHeadRepository = repositories.NewBusinessHeadRepository(DBConnection)
	businessHeadService    = services.NewBusinessHeadService(businessHeadRepository, logging)
	busHeadHandler         = NewBusinessHeadHandler(businessHeadService, logging, "Address")
)

func TestBusinessHeadHandler_GetAllBusinessHead(t *testing.T) {
	r := SetupRouter()
	r.GET("/v1/business_head", busHeadHandler.GetAllBusinessHead)

	request, err := http.NewRequest("GET", "/v1/business_head", nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var businessHead common.BusinessHeadDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &businessHead)
	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
}

func createBusinessHead(t *testing.T) *common.BusinessHeadDataResponse {
	r := SetupRouter()
	r.POST("/v1/business_head", busHeadHandler.CreateBusinessHead)
	entity := common.CreateBusinessHeadRequest{
		Company:                (&utils.Faker{}).RandomUUID(),
		JobTitle:               (&utils.Faker{}).RandomName(),
		Phone:                  (&utils.Faker{}).RandomString(11),
		IdentificationType:     (&utils.Faker{}).RandomString(20),
		IdentificationNumber:   (&utils.Faker{}).RandomString(15),
		IdentificationImageURL: (&utils.Faker{}).RandomString(15),
		CompanyIDUrl:           (&utils.Faker{}).RandomString(20),
	}

	jsonValue, _ := json.Marshal(entity)
	request, err := http.NewRequest("POST", "/v1/business_head", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error occurred decoding")
		return nil
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var businessHead *common.BusinessHeadDataResponse

	_ = json.Unmarshal([]byte(response.Body.String()), &businessHead)

	require.Equal(t, http.StatusCreated, response.Code)

	return businessHead
}

func TestBusinessHeadHandler_CreateBusinessHead(t *testing.T) {
	businessHead := createBusinessHead(t)
	require.NotEmpty(t, businessHead)
}

func TestBusinessHeadHandler_GetBusinessHeadByID(t *testing.T) {
	businessHead := createBusinessHead(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", businessHead.Data.ID.String())

	r.GET("/v1/business_head/:id", busHeadHandler.GetBusinessHeadByID)

	request, err := http.NewRequest("GET", "/v1/business_head/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.BusinessHeadDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, businessHead.Data.ID, resp.Data.ID)
	require.Equal(t, businessHead.Data.JobTitle, resp.Data.JobTitle)
	require.Equal(t, businessHead.Data.Company, resp.Data.Company)
	require.Equal(t, businessHead.Data.Phone, resp.Data.Phone)
	require.Equal(t, businessHead.Data.IdentificationType, resp.Data.IdentificationType)
	require.Equal(t, businessHead.Data.IdentificationNumber, resp.Data.IdentificationNumber)
	require.Equal(t, businessHead.Data.IdentificationImageURL, resp.Data.IdentificationImageURL)
	require.Equal(t, businessHead.Data.CompanyIDUrl, resp.Data.CompanyIDUrl)

}

func TestBusinessHeadHandler_GetBusinessHeadByCompanyID(t *testing.T) {
	businessHead := createBusinessHead(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", businessHead.Data.Company.String())

	r.GET("/v1/business_head/company/:id", busHeadHandler.GetBusinessHeadByCompanyID)

	request, err := http.NewRequest("GET", "/v1/business_head/company/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.BusinessHeadDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, businessHead.Data.ID, resp.Data.ID)
	require.Equal(t, businessHead.Data.JobTitle, resp.Data.JobTitle)
	require.Equal(t, businessHead.Data.Company, resp.Data.Company)
	require.Equal(t, businessHead.Data.Phone, resp.Data.Phone)
	require.Equal(t, businessHead.Data.IdentificationType, resp.Data.IdentificationType)
	require.Equal(t, businessHead.Data.IdentificationNumber, resp.Data.IdentificationNumber)
	require.Equal(t, businessHead.Data.IdentificationImageURL, resp.Data.IdentificationImageURL)
	require.Equal(t, businessHead.Data.CompanyIDUrl, resp.Data.CompanyIDUrl)

}

func TestBusinessHeadHandler_DeleteBusinessHead(t *testing.T) {
	r := SetupRouter()
	businessHead := createBusinessHead(t)

	r.DELETE("/v1/business_head/:id", busHeadHandler.DeleteBusinessHead)

	endpoint := fmt.Sprintf("/v1/business_head/%v", businessHead.Data.ID.String())

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestBusinessHeadHandler_UpdateBusinessHead(t *testing.T) {
	r := SetupRouter()
	businessHead := createBusinessHead(t)

	r.PATCH("/v1/business_head/:id", busHeadHandler.UpdateBusinessHead)

	endpoint := fmt.Sprintf("/v1/business_head/%v", businessHead.Data.ID.String())

	businessHeadJobTitle := (&utils.Faker{}).RandomName()
	businessHeadPhone := (&utils.Faker{}).RandomString(11)

	body := common.UpdateBusinessHeadRequest{
		JobTitle: &businessHeadJobTitle,
		Phone:    &businessHeadPhone,
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
