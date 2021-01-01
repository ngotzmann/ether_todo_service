package modules

import (
	"ether_todo/pkg/modules/config"
	"github.com/jinzhu/gorm"
	"strconv"
)


func DefaultGormDB(cfg config.Database) (*gorm.DB, error) {

	dsn := "host="+cfg.Address+
		   " port="+strconv.Itoa(cfg.Port)+
		   " user="+cfg.User+
		   " password="+cfg.Password+
		   " dbname="+cfg.Database+
		   " sslmode="+cfg.SSLMode

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(cfg.MaxIdleConnections)
	if cfg.ShouldLog {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	return db, nil
}

