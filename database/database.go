package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"
)

func New(config *Config, log *zap.Logger) (*gorm.DB, error) {
	logger := zapgorm2.New(log)
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
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                 logger,
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
