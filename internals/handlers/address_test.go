package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/soguazu/core_business/internals/common"
	"github.com/soguazu/core_business/internals/core/services"
	"github.com/soguazu/core_business/internals/repositories"
	"github.com/soguazu/core_business/pkg/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	addressRepository = repositories.NewAddressRepository(DBConnection)
	addressService    = services.NewAddressService(addressRepository, logging)
	addHandler        = NewAddressHandler(addressService, logging, "Address")
)

func TestAddressHandler_GetAllAddress(t *testing.T) {
	r := SetupRouter()
	r.GET("/v1/address", addHandler.GetAllAddress)

	request, err := http.NewRequest("GET", "/v1/address", nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var address common.GetAddressResponse

	err = json.Unmarshal([]byte(response.Body.String()), &address)
	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
}

func createAddress(t *testing.T) *common.CreateDataResponse {
	r := SetupRouter()
	r.POST("/v1/address", addHandler.CreateAddress)
	entity := common.CreateAddressRequest{
		Company:            (&utils.Faker{}).RandomUUID(),
		Address:            (&utils.Faker{}).RandomAddress(),
		City:               (&utils.Faker{}).RandomString(10),
		State:              (&utils.Faker{}).RandomString(10),
		Country:            (&utils.Faker{}).RandomString(10),
		UtilityBill:        (&utils.Faker{}).RandomString(15),
		ApartmentUnitFloor: (&utils.Faker{}).RandomInt(1, 10),
	}

	jsonValue, _ := json.Marshal(entity)
	request, err := http.NewRequest("POST", "/v1/address", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error occurred decoding")
		return nil
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var address *common.CreateDataResponse

	_ = json.Unmarshal([]byte(response.Body.String()), &address)

	require.Equal(t, http.StatusCreated, response.Code)

	return address
}

func TestAddressHandler_CreateAddress(t *testing.T) {
	address := createAddress(t)
	require.NotEmpty(t, address)

}

func TestAddressHandler_GetAddressByID(t *testing.T) {
	address := createAddress(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", address.Data.ID.String())

	r.GET("/v1/address/:id", addHandler.GetAddressByID)

	request, err := http.NewRequest("GET", "/v1/address/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.CreateDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, address.Data.ID, resp.Data.ID)
	require.Equal(t, address.Data.Address, resp.Data.Address)
	require.Equal(t, address.Data.Company, resp.Data.Company)
	require.Equal(t, address.Data.City, resp.Data.City)
	require.Equal(t, address.Data.State, resp.Data.State)
	require.Equal(t, address.Data.Country, resp.Data.Company)
	require.Equal(t, address.Data.UtilityBill, resp.Data.UtilityBill)
	require.Equal(t, address.Data.ApartmentUnitFloor, resp.Data.ApartmentUnitFloor)

}

func TestAddressHandler_GetAddressByCompanyID(t *testing.T) {
	address := createAddress(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", address.Data.Company.String())

	r.GET("/v1/address/company/:id", addHandler.GetAddressByID)

	request, err := http.NewRequest("GET", "/v1/address/company/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.CreateDataResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, address.Data.ID, resp.Data.ID)
	require.Equal(t, address.Data.Address, resp.Data.Address)
	require.Equal(t, address.Data.Company, resp.Data.Company)
	require.Equal(t, address.Data.City, resp.Data.City)
	require.Equal(t, address.Data.State, resp.Data.State)
	require.Equal(t, address.Data.Country, resp.Data.Company)
	require.Equal(t, address.Data.UtilityBill, resp.Data.UtilityBill)
	require.Equal(t, address.Data.ApartmentUnitFloor, resp.Data.ApartmentUnitFloor)
}

func TestAddressHandler_DeleteAddress(t *testing.T) {
	r := SetupRouter()
	address := createAddress(t)

	r.DELETE("/v1/address/:id", addHandler.DeleteAddress)

	endpoint := fmt.Sprintf("/v1/address/%v", address.Data.ID.String())

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestAddressHandler_UpdateAddress(t *testing.T) {
	r := SetupRouter()
	address := createAddress(t)

	r.PATCH("/v1/address/:id", addHandler.UpdateAddress)

	endpoint := fmt.Sprintf("/v1/address/%v", address.Data.ID.String())

	addressCountry := "Nigeria"
	addressCity := "Lagos"

	body := common.UpdateCompanyRequest{
		Name: &addressCountry,
		Type: &addressCity,
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
