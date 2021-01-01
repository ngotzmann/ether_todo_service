//+build wireinject

package modules

import (
	"ether_todo/pkg/modules/config"
	wire "github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func DefaultHttpServer() *echo.Echo {
	wire.Build(ProvideServiceCfg, DefaultEchoHttpServer)
	return &echo.Echo{}
}

func DefaultGorm() (*gorm.DB, error) {
	wire.Build(ProvideDBsCfg, DefaultGormDB)
	return &gorm.DB{}, nil
}

//Providers
func ProvideServiceCfg() config.Service {
	i := config.ReadConfig(config.Service{})
	cfg := i.(config.Service)
	return cfg
}

func ProvideDBsCfg() *config.Database {
	i := config.ReadConfig(config.Database{})
	cfg := i.(*config.Database)
	return cfg
}