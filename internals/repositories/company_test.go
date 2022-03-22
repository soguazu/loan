package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/soguazu/core_business/internals/common"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) *domain.Company {
	args := domain.Company{
		Owner:         (&utils.Faker{}).RandomUUID(),
		Name:          (&utils.Faker{}).RandomName(),
		Website:       (&utils.Faker{}).RandomWebsite(),
		Type:          (&utils.Faker{}).RandomType(),
		FundingSource: (&utils.Faker{}).RandomFundSource(),
		NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
	}

	company, err := CompanyRepository.Persist(&args)

	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, args.Name, company.Name)
	require.Equal(t, args.Owner, company.Owner)
	require.Equal(t, args.Website, company.Website)
	require.Equal(t, args.NoOfEmployee, company.NoOfEmployee)
	require.Equal(t, args.Type, company.Type)
	require.Equal(t, args.FundingSource, company.FundingSource)
	require.NotEmpty(t, company.ID)
	require.NotEmpty(t, company.CreatedAt)
	require.NotEmpty(t, company.UpdatedAt)

	return company
}

func TestPassedCreateCompany(t *testing.T) {
	for _, tc := range common.PassedTT {
		table := tc.Company
		t.Run(tc.TestName, func(t *testing.T) {
			c := domain.Company{
				Owner:         table.Owner,
				Name:          table.Name,
				Type:          table.Type,
				Website:       table.Website,
				FundingSource: table.FundingSource,
				NoOfEmployee:  table.NoOfEmployee,
			}

			company, err := CompanyRepository.Persist(&c)
			require.NoError(t, err)
			require.NotEmpty(t, company)

			require.Equal(t, c.Name, company.Name)
			require.Equal(t, c.Owner, company.Owner)
			require.Equal(t, c.Website, company.Website)
			require.Equal(t, c.NoOfEmployee, company.NoOfEmployee)
			require.Equal(t, c.Type, company.Type)
			require.Equal(t, c.FundingSource, company.FundingSource)
			require.NotEmpty(t, company.ID)
			require.NotEmpty(t, company.CreatedAt)
			require.NotEmpty(t, company.UpdatedAt)
		})

	}
}

func TestGetCompanyByID(t *testing.T) {
	randomCompany := createRandomAccount(t)
	company, err := CompanyRepository.GetByID(randomCompany.ID.String())
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, company.Owner, randomCompany.Owner)
	require.Equal(t, company.Type, randomCompany.Type)
	require.Equal(t, company.Website, randomCompany.Website)
	require.Equal(t, company.FundingSource, randomCompany.FundingSource)
	require.WithinDuration(t, company.CreatedAt, randomCompany.CreatedAt, time.Second)
}

func TestGetCompanyByBadID(t *testing.T) {
	company, err := CompanyRepository.GetByID(uuid.NewV4().String())
	require.Error(t, err)
	require.Empty(t, company)
}

func TestUpdateCompany(t *testing.T) {
	randomCompany := createRandomAccount(t)
	randomCompany.Name = "GTB2"
	company, err := CompanyRepository.Persist(randomCompany)
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, company.Name, randomCompany.Name)
}

func TestDeleteCompany(t *testing.T) {
	randomCompany := createRandomAccount(t)
	err := CompanyRepository.Delete(randomCompany.ID.String())
	require.NoError(t, err)

	company, err := CompanyRepository.GetByID(randomCompany.ID.String())
	require.Error(t, err)
	require.Empty(t, company)
}

func TestGetAllCompany(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	pagination := &utils.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "created_at asc",
	}
	p, err := CompanyRepository.Get(pagination)
	require.NoError(t, err)
	require.Len(t, p.Rows, 2)
}

func TestDeleteAll(t *testing.T) {
	t.Cleanup(func() {
		err := CompanyRepository.DeleteAll()
		if err != nil {
			return
		}
	})
}
