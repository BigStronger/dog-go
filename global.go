package dog

import (
	"github.com/BigStronger/dog-go/cache"
	"github.com/BigStronger/dog-go/database"
	"github.com/BigStronger/dog-go/id"
	"github.com/BigStronger/dog-go/log"
	"github.com/BigStronger/dog-go/token"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log   *zap.Logger
	DB    *gorm.DB
	Cache cache.API
	IDGen id.API
	Token token.API
)

func InitLog(config *log.Config) error {
	_log, err := log.New(config)
	if err == nil {
		Log = _log
	}
	return err
}

func InitDB(config *database.Config) error {
	_database, err := database.New(config)
	if err == nil {
		DB = _database
	}
	return err
}

func InitCache(config *cache.Config) error {
	_cache, err := cache.New(config)
	if err == nil {
		Cache = _cache
	}
	return err
}

func InitIDGenWithCache(cacheAPI cache.API, prefix string) {
	IDGen = id.NewWithCache(cacheAPI, prefix)
}
func InitIDGenWithLocal(nodeId int64) {
	IDGen = id.NewWithLocal(nodeId)
}

func InitToken(config *token.Config) {
	Token = token.New(config)
}
