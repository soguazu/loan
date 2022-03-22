package repositories

import (
	log "github.com/sirupsen/logrus"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/database"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

var (
	DBConnection      *gorm.DB
	CompanyRepository ports.ICompanyRepository
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
}
