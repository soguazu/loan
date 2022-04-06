package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/database"
	"core_business/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

var (
	DBConnection              *gorm.DB
	CompanyRepository         ports.ICompanyRepository
	AddressRepository         ports.IAddressRepository
	BusinessHeadRepository    ports.IBusinessHeadRepository
	BusinessPartnerRepository ports.IBusinessPartnerRepository
	CompanyProfileRepository  ports.ICompanyProfileRepository
	Company                   *domain.Company
)

func TestMain(m *testing.M) {
	db := database.NewSqliteDatabase()

	DBConnection = db.ConnectDB(filepath.Join("..", "..", "evea.db"))
	err := db.MigrateAll(DBConnection)
	if err != nil {
		log.Fatal(err)
	}

	instantiateRepos()

	os.Exit(m.Run())

}

func instantiateRepos() {
	CompanyRepository = &companyRepository{
		db: DBConnection,
	}
	Company = &domain.Company{
		Owner:         (&utils.Faker{}).RandomUUID(),
		Name:          (&utils.Faker{}).RandomName(),
		Website:       (&utils.Faker{}).RandomWebsite(),
		Type:          (&utils.Faker{}).RandomType(),
		FundingSource: (&utils.Faker{}).RandomFundSource(),
		NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
	}

	CompanyRepository.Persist(Company)

	AddressRepository = &addressRepository{
		db: DBConnection,
	}

	BusinessHeadRepository = &businessHeadRepository{
		db: DBConnection,
	}

	BusinessPartnerRepository = &businessPartnerRepository{
		db: DBConnection,
	}

	CompanyProfileRepository = &companyProfileRepository{
		db: DBConnection,
	}
}
