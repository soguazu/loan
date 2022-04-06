package repositories

import (
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomCompanyProfile(t *testing.T) *domain.CompanyProfile {
	args := &domain.CompanyProfile{
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

	err := CompanyProfileRepository.Persist(args)
	require.NoError(t, err)

	if err != nil {
		return nil
	}

	companyProfile, err := CompanyProfileRepository.GetByID(args.ID.String())

	require.NotEmpty(t, companyProfile)
	require.NoError(t, err)

	require.Equal(t, args.RCNumber, companyProfile.RCNumber)
	require.Equal(t, args.Company, companyProfile.Company)
	require.Equal(t, args.BusinessTin, companyProfile.BusinessTin)
	require.Equal(t, args.BusinessType, companyProfile.BusinessType)
	require.Equal(t, args.BusinessActivity, companyProfile.BusinessActivity)
	require.Equal(t, args.CACCertificateURL, companyProfile.CACCertificateURL)
	require.Equal(t, args.MermatURL, companyProfile.MermatURL)
	require.Equal(t, args.StatusReportURL, companyProfile.StatusReportURL)

	require.NotEmpty(t, companyProfile.ID)
	require.NotEmpty(t, companyProfile.CreatedAt)
	require.NotEmpty(t, companyProfile.UpdatedAt)

	return args
}

func TestCompanyProfileRepository_Persist(t *testing.T) {
	randomCompanyProfile := createRandomCompanyProfile(t)
	companyProfile, err := CompanyProfileRepository.GetByID(randomCompanyProfile.ID.String())

	require.NoError(t, err)
	require.NotEmpty(t, companyProfile)
	require.Equal(t, randomCompanyProfile.Company, companyProfile.Company)
	require.Equal(t, randomCompanyProfile.BusinessTin, companyProfile.BusinessTin)
	require.Equal(t, randomCompanyProfile.BusinessType, companyProfile.BusinessType)
	require.Equal(t, randomCompanyProfile.BusinessActivity, companyProfile.BusinessActivity)
	require.Equal(t, randomCompanyProfile.CACCertificateURL, companyProfile.CACCertificateURL)
	require.Equal(t, randomCompanyProfile.MermatURL, companyProfile.MermatURL)
	require.Equal(t, randomCompanyProfile.StatusReportURL, companyProfile.StatusReportURL)

	require.WithinDuration(t, companyProfile.CreatedAt, randomCompanyProfile.CreatedAt, time.Second)
}

func TestGetCompanyProfileByBadID(t *testing.T) {
	companyProfile, err := CompanyProfileRepository.GetByID(uuid.NewV4().String())
	require.Error(t, err)
	require.Empty(t, companyProfile)
}

func TestCompanyProfileRepository_Update(t *testing.T) {
	randomCompanyProfile := createRandomCompanyProfile(t)
	randomCompanyProfile.IncorporationYear = time.Now().String()

	err := CompanyProfileRepository.Persist(randomCompanyProfile)

	companyProfile, err := CompanyProfileRepository.GetByID(randomCompanyProfile.ID.String())
	if err != nil {
		return
	}

	require.NoError(t, err)
	require.NotEmpty(t, companyProfile)
	require.Equal(t, companyProfile.IncorporationYear, randomCompanyProfile.IncorporationYear)
}

func TestCompanyProfileRepository_Delete(t *testing.T) {
	randomCompanyProfile := createRandomCompanyProfile(t)
	err := CompanyProfileRepository.Delete(randomCompanyProfile.ID.String())
	require.NoError(t, err)

	companyProfile, err := CompanyProfileRepository.GetByID(randomCompanyProfile.ID.String())
	require.Error(t, err)
	require.Empty(t, companyProfile)
}

func TestCompanyProfileRepository_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCompanyProfile(t)
	}

	pagination := &utils.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "created_at asc",
	}
	p, err := CompanyProfileRepository.Get(pagination)
	require.NoError(t, err)
	require.Len(t, p.Rows, 2)
}

func TestDeleteAllCompanyProfile(t *testing.T) {
	t.Cleanup(func() {
		err := CompanyProfileRepository.DeleteAll()
		if err != nil {
			return
		}
	})
}
