package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomBusinessPartner(t *testing.T) *domain.BusinessPartner {
	args := &domain.BusinessPartner{
		Company: (&utils.Faker{}).RandomUUID(),
		Name:    (&utils.Faker{}).RandomName(),
		Phone:   (&utils.Faker{}).RandomString(11),
	}

	err := BusinessPartnerRepository.Persist(args)
	require.NoError(t, err)

	if err != nil {
		return nil
	}

	businessPartner, err := BusinessPartnerRepository.GetByID(args.ID.String())

	require.NotEmpty(t, businessPartner)
	require.NoError(t, err)

	require.Equal(t, args.Name, businessPartner.Name)
	require.Equal(t, args.Company, businessPartner.Company)
	require.Equal(t, args.Phone, businessPartner.Phone)

	require.NotEmpty(t, businessPartner.ID)
	require.NotEmpty(t, businessPartner.CreatedAt)
	require.NotEmpty(t, businessPartner.UpdatedAt)

	return args
}

func TestBusinessPartnerRepository_GetByID(t *testing.T) {
	randomBusinessPartner := createRandomBusinessPartner(t)
	businessPartner, err := BusinessPartnerRepository.GetByID(randomBusinessPartner.ID.String())

	require.NoError(t, err)
	require.NotEmpty(t, businessPartner)
	require.Equal(t, randomBusinessPartner.Company, businessPartner.Company)
	require.Equal(t, randomBusinessPartner.Name, businessPartner.Name)
	require.Equal(t, randomBusinessPartner.Phone, businessPartner.Phone)

	require.WithinDuration(t, businessPartner.CreatedAt, randomBusinessPartner.CreatedAt, time.Second)
}

func TestGetBusinessPartnerByBadID(t *testing.T) {
	businessPartner, err := BusinessPartnerRepository.GetByID(uuid.NewV4().String())
	require.Error(t, err)
	require.Empty(t, businessPartner)
}

func TestBusinessPartnerRepository_Update(t *testing.T) {
	randomBusinessPartner := createRandomBusinessPartner(t)
	randomBusinessPartner.Name = "John Doe"

	err := BusinessPartnerRepository.Persist(randomBusinessPartner)

	businessPartner, err := BusinessPartnerRepository.GetByID(randomBusinessPartner.ID.String())
	if err != nil {
		return
	}

	require.NoError(t, err)
	require.NotEmpty(t, businessPartner)
	require.Equal(t, businessPartner.Name, randomBusinessPartner.Name)
}

func TestBusinessPartnerRepository_Delete(t *testing.T) {
	randomBusinessPartner := createRandomBusinessPartner(t)
	err := BusinessPartnerRepository.Delete(randomBusinessPartner.ID.String())
	require.NoError(t, err)

	businessPartner, err := BusinessHeadRepository.GetByID(randomBusinessPartner.ID.String())
	require.Error(t, err)
	require.Empty(t, businessPartner)
}

func TestBusinessPartnerRepository_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBusinessPartner(t)
	}

	pagination := &utils.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "created_at asc",
	}
	p, err := BusinessPartnerRepository.Get(pagination)
	require.NoError(t, err)
	require.Len(t, p.Rows, 2)
}

func TestDeleteAllBusinessPartner(t *testing.T) {
	t.Cleanup(func() {
		err := BusinessPartnerRepository.DeleteAll()
		if err != nil {
			return
		}
	})
}
