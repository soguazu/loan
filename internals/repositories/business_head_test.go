package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomBusinessHead(t *testing.T) *domain.BusinessHead {
	args := &domain.BusinessHead{
		Company:                (&utils.Faker{}).RandomUUID(),
		JobTitle:               (&utils.Faker{}).RandomName(),
		Phone:                  (&utils.Faker{}).RandomString(11),
		IdentificationType:     (&utils.Faker{}).RandomString(20),
		IdentificationNumber:   (&utils.Faker{}).RandomString(15),
		IdentificationImageURL: (&utils.Faker{}).RandomString(15),
		CompanyIDUrl:           (&utils.Faker{}).RandomString(20),
	}

	err := BusinessHeadRepository.Persist(args)
	require.NoError(t, err)

	if err != nil {
		return nil
	}

	businessHead, err := BusinessHeadRepository.GetByID(args.ID.String())

	require.NotEmpty(t, businessHead)
	require.NoError(t, err)

	require.Equal(t, args.JobTitle, businessHead.JobTitle)
	require.Equal(t, args.Company, businessHead.Company)
	require.Equal(t, args.Phone, businessHead.Phone)
	require.Equal(t, args.IdentificationType, businessHead.IdentificationType)
	require.Equal(t, args.IdentificationNumber, businessHead.IdentificationNumber)
	require.Equal(t, args.IdentificationImageURL, businessHead.IdentificationImageURL)
	require.Equal(t, args.CompanyIDUrl, businessHead.CompanyIDUrl)

	require.NotEmpty(t, businessHead.ID)
	require.NotEmpty(t, businessHead.CreatedAt)
	require.NotEmpty(t, businessHead.UpdatedAt)

	return args
}

func TestBusinessHeadRepository_GetByID(t *testing.T) {
	randomBusinessHead := createRandomBusinessHead(t)
	businessHead, err := BusinessHeadRepository.GetByID(randomBusinessHead.ID.String())

	require.NoError(t, err)
	require.NotEmpty(t, businessHead)
	require.Equal(t, randomBusinessHead.Company, businessHead.Company)
	require.Equal(t, randomBusinessHead.JobTitle, businessHead.JobTitle)
	require.Equal(t, randomBusinessHead.Phone, businessHead.Phone)
	require.Equal(t, randomBusinessHead.IdentificationType, businessHead.IdentificationType)
	require.Equal(t, randomBusinessHead.IdentificationNumber, businessHead.IdentificationNumber)
	require.Equal(t, randomBusinessHead.IdentificationImageURL, businessHead.IdentificationImageURL)
	require.Equal(t, randomBusinessHead.CompanyIDUrl, businessHead.CompanyIDUrl)
	require.WithinDuration(t, businessHead.CreatedAt, randomBusinessHead.CreatedAt, time.Second)
}

func TestGetBusinessHeadByBadID(t *testing.T) {
	address, err := BusinessHeadRepository.GetByID(uuid.NewV4().String())
	require.Error(t, err)
	require.Empty(t, address)
}

func TestUpdateBusinessHead(t *testing.T) {
	randomBusinessHead := createRandomBusinessHead(t)
	randomBusinessHead.JobTitle = "ceo"

	err := BusinessHeadRepository.Persist(randomBusinessHead)

	businessHead, err := BusinessHeadRepository.GetByID(randomBusinessHead.ID.String())
	if err != nil {
		return
	}

	require.NoError(t, err)
	require.NotEmpty(t, businessHead)
	require.Equal(t, businessHead.JobTitle, randomBusinessHead.JobTitle)
}

func TestBusinessHeadRepository_Delete(t *testing.T) {
	randomBusinessHead := createRandomBusinessHead(t)
	err := BusinessHeadRepository.Delete(randomBusinessHead.ID.String())
	require.NoError(t, err)

	businessHead, err := BusinessHeadRepository.GetByID(randomBusinessHead.ID.String())
	require.Error(t, err)
	require.Empty(t, businessHead)
}

func TestBusinessHeadRepository_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBusinessHead(t)
	}

	pagination := &utils.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "created_at asc",
	}
	p, err := BusinessHeadRepository.Get(pagination)
	require.NoError(t, err)
	require.Len(t, p.Rows, 2)
}

func TestDeleteAllBusinessHead(t *testing.T) {
	t.Cleanup(func() {
		err := BusinessHeadRepository.DeleteAll()
		if err != nil {
			return
		}
	})
}
