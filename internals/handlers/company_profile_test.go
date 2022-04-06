package handlers

import (
	"bytes"
	"core_business/internals/common"
	"core_business/internals/core/services"
	"core_business/internals/repositories"
	"core_business/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

var (
	companyProfileRepository = repositories.NewCompanyProfileRepository(DBConnection)
	companyProfileService    = services.NewCompanyProfileService(companyProfileRepository, logging)
	comProfileHandler        = NewCompanyProfileHandler(companyProfileService, logging, "Company profile")
)

func TestCompanyProfileHandler_GetAllCompanyProfile(t *testing.T) {
	r := SetupRouter()
	r.GET("/v1/company_profile", comProfileHandler.GetAllCompanyProfile)

	request, err := http.NewRequest("GET", "/v1/company_profile", nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var companyProfile common.CompanyProfileDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &companyProfile)
	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
}

func createCompanyProfile(t *testing.T) *common.CompanyProfileDataResponse {
	r := SetupRouter()
	r.POST("/v1/company_profile", comProfileHandler.CreateCompanyProfile)
	entity := common.CreateCompanyProfileRequest{
		Company:            (&utils.Faker{}).RandomUUID(),
		RCNumber:           (&utils.Faker{}).RandomString(15),
		BusinessTin:        (&utils.Faker{}).RandomString(15),
		BusinessType:       (&utils.Faker{}).RandomString(15),
		BusinessActivity:   (&utils.Faker{}).RandomString(15),
		IncorporationYear:  time.Now().String(),
		IncorporationState: (&utils.Faker{}).RandomString(10),
		CACCertificateURL:  (&utils.Faker{}).RandomString(15),
		MermatURL:          (&utils.Faker{}).RandomString(15),
		StatusReportURL:    (&utils.Faker{}).RandomString(15),
	}

	jsonValue, _ := json.Marshal(entity)
	request, err := http.NewRequest("POST", "/v1/company_profile", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error occurred decoding")
		return nil
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var companyProfile *common.CompanyProfileDataResponse

	_ = json.Unmarshal([]byte(response.Body.String()), &companyProfile)

	require.Equal(t, http.StatusCreated, response.Code)

	return companyProfile
}

func TestCompanyProfileHandler_CreateCompanyProfile(t *testing.T) {
	companyProfile := createCompanyProfile(t)
	require.NotEmpty(t, companyProfile)
}

func TestCompanyProfileHandler_GetCompanyProfileByID(t *testing.T) {
	companyProfile := createCompanyProfile(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", companyProfile.Data.ID.String())

	r.GET("/v1/business_partner/:id", busPartnerHandler.GetBusinessPartnerByID)

	request, err := http.NewRequest("GET", "/v1/business_partner/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.CompanyProfileDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, companyProfile.Data.ID, resp.Data.ID)
	require.Equal(t, companyProfile.Data.RCNumber, resp.Data.RCNumber)
	require.Equal(t, companyProfile.Data.Company, resp.Data.Company)
	require.Equal(t, companyProfile.Data.BusinessTin, resp.Data.BusinessTin)
	require.Equal(t, companyProfile.Data.BusinessType, resp.Data.BusinessType)
	require.Equal(t, companyProfile.Data.BusinessActivity, resp.Data.BusinessActivity)
	require.Equal(t, companyProfile.Data.IncorporationState, resp.Data.IncorporationState)
	require.Equal(t, companyProfile.Data.IncorporationYear, resp.Data.IncorporationYear)
	require.Equal(t, companyProfile.Data.CACCertificateURL, resp.Data.CACCertificateURL)
	require.Equal(t, companyProfile.Data.MermatURL, resp.Data.MermatURL)
	require.Equal(t, companyProfile.Data.StatusReportURL, resp.Data.StatusReportURL)
}

func TestCompanyProfileHandler_GetCompanyProfileByCompanyID(t *testing.T) {
	companyProfile := createCompanyProfile(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", companyProfile.Data.Company.String())

	r.GET("/v1/company_profile/company/:id", comProfileHandler.GetCompanyProfileByCompanyID)

	request, err := http.NewRequest("GET", "/v1/company_profile/company/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.CompanyProfileDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, companyProfile.Data.ID, resp.Data.ID)
	require.Equal(t, companyProfile.Data.RCNumber, resp.Data.RCNumber)
	require.Equal(t, companyProfile.Data.Company, resp.Data.Company)
	require.Equal(t, companyProfile.Data.BusinessTin, resp.Data.BusinessTin)
	require.Equal(t, companyProfile.Data.BusinessType, resp.Data.BusinessType)
	require.Equal(t, companyProfile.Data.BusinessActivity, resp.Data.BusinessActivity)
	require.Equal(t, companyProfile.Data.IncorporationState, resp.Data.IncorporationState)
	require.Equal(t, companyProfile.Data.IncorporationYear, resp.Data.IncorporationYear)
	require.Equal(t, companyProfile.Data.CACCertificateURL, resp.Data.CACCertificateURL)
	require.Equal(t, companyProfile.Data.MermatURL, resp.Data.MermatURL)
	require.Equal(t, companyProfile.Data.StatusReportURL, resp.Data.StatusReportURL)
}

func TestCompanyProfileHandler_DeleteCompanyProfile(t *testing.T) {
	r := SetupRouter()
	companyProfile := createCompanyProfile(t)

	r.DELETE("/v1/company_profile/:id", comProfileHandler.DeleteCompanyProfile)

	endpoint := fmt.Sprintf("/v1/company_profile/%v", companyProfile.Data.ID.String())

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestCompanyProfileHandler_UpdateCompanyProfile(t *testing.T) {
	r := SetupRouter()
	companyProfile := createCompanyProfile(t)

	r.PATCH("/v1/company_profile/:id", comProfileHandler.UpdateCompanyProfile)

	endpoint := fmt.Sprintf("/v1/company_profile/%v", companyProfile.Data.ID.String())
	companyProfileBusinessType := (&utils.Faker{}).RandomName()
	companyProfileBusinessActivity := (&utils.Faker{}).RandomString(11)

	body := common.UpdateCompanyProfileRequest{
		BusinessType:     &companyProfileBusinessType,
		BusinessActivity: &companyProfileBusinessActivity,
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
