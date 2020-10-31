//+build wireinject

package modules

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func DefaultHttpServer() *echo.Echo {
	wire.Build(DefaultConfig, DefaultEchoHttpServer)
	return &echo.Echo{}
}

func DefaultGorm() (*gorm.DB, error) {
	wire.Build(DefaultConfig, DefaultFileLogger, DefaultGormDB)
	return &gorm.DB{}, nil
}
