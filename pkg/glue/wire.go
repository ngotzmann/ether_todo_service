//+build wireinject

package glue

import (
	"ether_todo/pkg/glue/config"
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
func ProvideServiceCfg() config.Server {
	i := config.ReadConfig(config.Server{})
	cfg := i.(config.Server)
	return cfg
}

func ProvideDBsCfg() config.Database {
	i := config.ReadConfig(config.Database{})
	cfg := i.(config.Database)
	return cfg
}