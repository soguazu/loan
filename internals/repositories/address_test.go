package repositories

import (
	uuid "github.com/satori/go.uuid"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAddress(t *testing.T) *domain.Address {
	args := &domain.Address{
		Company:            (&utils.Faker{}).RandomUUID(),
		Address:            (&utils.Faker{}).RandomAddress(),
		City:               (&utils.Faker{}).RandomString(10),
		State:              (&utils.Faker{}).RandomString(10),
		Country:            (&utils.Faker{}).RandomString(10),
		UtilityBill:        (&utils.Faker{}).RandomString(15),
		ApartmentUnitFloor: (&utils.Faker{}).RandomInt(1, 10),
	}

	err := AddressRepository.Persist(args)
	require.NoError(t, err)

	if err != nil {
		return nil
	}

	address, err := AddressRepository.GetByID(args.ID.String())

	require.NotEmpty(t, address)
	require.NoError(t, err)

	require.Equal(t, args.Address, address.Address)
	require.Equal(t, args.Company, address.Company)
	require.Equal(t, args.City, address.City)
	require.Equal(t, args.State, address.State)
	require.Equal(t, args.Country, address.Country)

	require.NotEmpty(t, address.ID)
	require.NotEmpty(t, address.CreatedAt)
	require.NotEmpty(t, address.UpdatedAt)

	return args
}

func TestGetAddressByID(t *testing.T) {
	randomAddress := createRandomAddress(t)
	address, err := AddressRepository.GetByID(randomAddress.ID.String())

	require.NoError(t, err)
	require.NotEmpty(t, address)
	require.Equal(t, randomAddress.Address, address.Address)
	require.Equal(t, randomAddress.Company, address.Company)
	require.Equal(t, randomAddress.City, address.City)
	require.Equal(t, randomAddress.State, address.State)
	require.Equal(t, randomAddress.Country, address.Country)
	require.WithinDuration(t, address.CreatedAt, randomAddress.CreatedAt, time.Second)
}

func TestGetAddressByBadID(t *testing.T) {
	address, err := AddressRepository.GetByID(uuid.NewV4().String())
	require.Error(t, err)
	require.Empty(t, address)
}

func TestUpdateAddress(t *testing.T) {
	randomAddress := createRandomAddress(t)
	randomAddress.Address = "No 1 kafaye street"

	err := AddressRepository.Persist(randomAddress)

	address, err := AddressRepository.GetByID(randomAddress.ID.String())
	if err != nil {
		return
	}

	require.NoError(t, err)
	require.NotEmpty(t, address)
	require.Equal(t, address.Address, randomAddress.Address)
}

func TestDeleteAddress(t *testing.T) {
	randomAddress := createRandomAddress(t)
	err := AddressRepository.Delete(randomAddress.ID.String())
	require.NoError(t, err)

	address, err := AddressRepository.GetByID(randomAddress.ID.String())
	require.Error(t, err)
	require.Empty(t, address)
}

func TestGetAllAddress(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAddress(t)
	}

	pagination := &utils.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "created_at asc",
	}
	p, err := AddressRepository.Get(pagination)
	require.NoError(t, err)
	require.Len(t, p.Rows, 2)
}

func TestDeleteAllAddress(t *testing.T) {
	t.Cleanup(func() {
		err := AddressRepository.DeleteAll()
		if err != nil {
			return
		}
	})
}
