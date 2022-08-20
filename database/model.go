package database

import "gorm.io/gorm"

type Config struct {
	Addr        string `yaml:"addr"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Config      string `yaml:"config"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxLifetime int64  `yaml:"maxLifetime"`
}

type CommitCallback func(tx *gorm.DB) error

type Paging struct {
	Page int `form:"page" json:"page"`
	Rows int `form:"rows" json:"rows"`
}
