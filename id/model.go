package id

import (
	"github.com/bwmarrin/snowflake"
)

type API interface {
	Generate() snowflake.ID
	UUID() string
	UUID32() string
}
