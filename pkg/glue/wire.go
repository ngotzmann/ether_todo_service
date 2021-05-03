//+build wireinject

package glue

import (
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
func ProvideServiceCfg() Server {
	i := ReadConfig(Server{})
	cfg := i.(Server)
	return cfg
}

func ProvideDBsCfg() Database {
	i := ReadConfig(Database{})
	cfg := i.(Database)
	return cfg
}

func ProvideLogCfg() Log {
	i := ReadConfig(Log{})
	cfg := i.(Log)
	return cfg
}