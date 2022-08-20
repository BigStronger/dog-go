package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func New(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v",
		config.Username,
		config.Password,
		config.Addr,
		config.Database,
		config.Config,
	)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
	}
	var dbLog logger.Interface
	dbLog = logger.New(log.New(os.Stdout, "", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                 dbLog,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxLifetime) * time.Second)
	return db, nil
}
