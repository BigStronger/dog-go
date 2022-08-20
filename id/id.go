package id

import (
	"fmt"
	"github.com/BigStronger/dog-go/cache"
	"github.com/bwmarrin/snowflake"
	"os"
	"time"
)

type _IDGen struct {
	cache  cache.API
	prefix *string
	node   *snowflake.Node
}

func (curr *_IDGen) Generate() snowflake.ID {
	return curr.node.Generate()
}

func NewWithCache(cache cache.API, prefix string) API {
	r := _IDGen{cache: cache, prefix: &prefix}
	r.loadSnowflakeNodeID()
	return &r
}

func NewWithLocal(nodeId int64) API {
	node, _ := snowflake.NewNode(nodeId)
	return &_IDGen{node: node}
}

func (curr *_IDGen) loadSnowflakeNodeID() {
	nodeId := int64(0)
	for {
		if nodeId > 1023 {
			os.Exit(1)
		}
		if success, err := curr.cache.SetNX(curr.cache.CacheTimeoutContext(), fmt.Sprintf("%s:%d", *curr.prefix, nodeId), true, time.Second*30).Result(); err == nil && success {
			break
		}
		nodeId++
	}
	node, _ := snowflake.NewNode(nodeId)
	curr.node = node
	go curr.snowflakeNodeIDKeepalive(nodeId)
}

func (curr *_IDGen) snowflakeNodeIDKeepalive(nodeId int64) {
	tryNum := 0
	for {
		if tryNum > 20 {
			os.Exit(1)
		}
		if success, err := curr.cache.Expire(curr.cache.CacheTimeoutContextWithDuration(time.Second), fmt.Sprintf("%s:%d", *curr.prefix, nodeId), time.Second*30).Result(); err != nil || !success {
			tryNum++
			time.Sleep(time.Millisecond * 500)
			continue
		}
		tryNum = 0
		time.Sleep(time.Second * 5)
	}
}
