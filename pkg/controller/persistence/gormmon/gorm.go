package gormmon

import (
	"github.com/jinzhu/gorm"
	"github.com/ngotzmann/gorror"
	loggable "github.com/sas1024/gorm-loggable"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

var db *gorm.DB

type GormConfig struct {
	Host string
	Port string
	DBName string
	Username string
	Password string
	MaxIdleConnections int
	ShouldLog bool
	ShouldAudit bool
}

func GetGormDB() (*gorm.DB, error) {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		err := gorror.CreateError(gorror.DatabaseError, "Database connection is not initialised")
		return nil, err
	} else {
		return db, nil
	}
}

func InitGormDB(c GormConfig) {
	gdb, err := gorm.Open(
		"postgres",
		"host="+c.Host+
			" port="+c.Port+
			" user="+c.Username+
			" dbname="+c.DBName+
			" password="+c.Password)

	if err != nil {
		log.Panic("Database opening error: ", err)
	}
	gdb.DB().SetMaxIdleConns(c.MaxIdleConnections)
	if c.ShouldLog {
		//db.SetLogger(&log.)
		gdb.LogMode(true)
	} else {
		gdb.LogMode(false)
	}

	if !true {
		_, err = loggable.Register(gdb)
		if err != nil {
			log.Panic("Loggable plugin can not initialized: ", err)
		}
	}

	db = gdb
}
