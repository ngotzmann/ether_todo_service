package modules

import (
	"github.com/jinzhu/gorm"
)


func DefaultGormDB(cfg *Config, log *Logger) (*gorm.DB, error) {

	dsn := "host="+cfg.DBAddress+
		   " port="+cfg.DBPort+
		   " user="+cfg.DBUser+
		   " password="+cfg.DBPassword+
		   " dbname="+cfg.Database+
		   " sslmode="+cfg.DBSSLMode

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(cfg.DBMaxIdleConnections)
	if cfg.DBShouldLog {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	return db, nil
}

