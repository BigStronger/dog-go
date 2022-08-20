package id

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

func UUID() string {
	now := time.Now().Format("200601021504")
	id, _ := strconv.ParseUint(strings.ReplaceAll(uuid.NewV4().String(), "-", "")[:16], 16, 64)
	uid := fmt.Sprintf("%s%020d", now, id)
	return uid
}

func UUID32() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func ParseString(id string) (snowflake.ID, error) {
	return snowflake.ParseString(id)
}
