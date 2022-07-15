package manager

import (
	"fmt"
	"golang-upload-download/config"
	"golang-upload-download/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Infra disini bertugas sebagai database penyimpanan pengganti slice
type InfraManager interface {
	SqlDb() *gorm.DB
	FilePath() string
}

type infraManager struct {
	db  *gorm.DB
	cfg config.Config
}

func (i *infraManager) SqlDb() *gorm.DB {
	return i.db
}

func (i *infraManager) FilePath() string {
	return i.cfg.FilePath
}

func (i *infraManager) initDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", i.cfg.Host, i.cfg.User, i.cfg.Password, i.cfg.DbName, i.cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	i.db = db
	env := os.Getenv("ENV")
	if env == "migration" {
		db.Debug()
		db.AutoMigrate(&model.Product{})
	} else if env == "dev" {
		db.Debug()
	}
}

func NewInfra(config config.Config) InfraManager {
	infra := infraManager{cfg: config}
	infra.initDb()
	return &infra
}
