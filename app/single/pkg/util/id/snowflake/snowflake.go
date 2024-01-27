package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

type Snowflake struct {
	node *snowflake.Node
}

func NewFlake(n int64) *Snowflake {
	sf, err := snowflake.NewNode(n)
	if err != nil {
		log.Fatal(err)
	}
	return &Snowflake{
		sf,
	}
}

// NewSnowflake 生成雪花算法id
func (s *Snowflake) Generate() int64 {
	return s.node.Generate().Int64()
}

// NewSnowflake 生成雪花算法id
func (s *Snowflake) Parse(id int64) snowflake.ID {
	return snowflake.ID(id)
}
