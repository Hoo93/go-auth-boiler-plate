package data

import (
	"auth-server-boiler-plate/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
// wire 에 등록해 의존성이 자동으로 관리됨
var ProviderSet = wire.NewSet(
	NewData,
	NewGormClient,
	NewGreeterRepo,
	NewUserRepository,
)

// Data .
type Data struct {
	log *log.Helper
	db  *gorm.DB
}

// NewData .
//func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	cleanup := func() {
//		log.NewHelper(logger).Info("closing the data resources")
//	}
//	return &Data{}, cleanup, nil
//}

func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	log := log.NewHelper(logger)
	log.Info("initializing data layer")

	cleanup := func() {
		log.Info("closing the data resources")
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return &Data{
		log: log,
		db:  db,
	}, cleanup, nil
}

func NewGormClient(cfg *conf.Data, logger log.Logger) *gorm.DB {
	log.NewHelper(logger).Info("init gorm client")

	driver := postgres.Open(cfg.Database.Source)

	client, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Errorf("failed to connect to database: %v", err)
		return nil
	}

	return client
}
