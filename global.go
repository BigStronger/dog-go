package dog

import (
	"github.com/BigStronger/dog-go/cache"
	"github.com/BigStronger/dog-go/database"
	"github.com/BigStronger/dog-go/id"
	"github.com/BigStronger/dog-go/log"
	"github.com/BigStronger/dog-go/token"
	"gorm.io/gorm"
)

var (
	Log      log.API
	Database *gorm.DB
	Cache    cache.API
	ID       id.API
	Token    token.API
)

func InitLog(config *log.Config) error {
	_log, err := log.New(config)
	if err == nil {
		Log = _log
	}
	return err
}

func InitDatabase(config *database.Config) error {
	_database, err := database.New(config)
	if err == nil {
		Database = _database
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

func InitIDWithCache(cacheAPI cache.API, prefix string) {
	ID = id.NewWithCache(cacheAPI, prefix)
}
func InitIDWithLocal(nodeId int64) {
	ID = id.NewWithLocal(nodeId)
}

func InitToken(config *token.Config) {
	Token = token.New(config)
}
