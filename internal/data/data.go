package data

import (
	"auth-server-boiler-plate/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	log *log.Helper
	db  *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewGormClient(cfg *conf.Bootstrap, logger log.Logger) *gorm.DB {
	// TODO initialize gorm.DB client
	log.NewHelper(logger).Info("init gorm client")

	driver := postgres.Open(cfg.Data.Database.Source)

	client, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Errorf("failed to connect to database: %v", err)
		return nil
	}

	return client
}
