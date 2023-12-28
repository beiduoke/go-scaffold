package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

type Snowflake struct {
	node *snowflake.Node
}

func NewFlake() *Snowflake {
	sf, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	return &Snowflake{
		sf,
	}
}

// NewSnowflake 生成雪花算法id
func (s *Snowflake) Generate(logger log.Logger) int64 {
	return s.node.Generate().Int64()
}

// NewSnowflake 生成雪花算法id
func (s *Snowflake) Parse(logger log.Logger) *snowflake.Node {
	sf, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal("snowflake no init")
	}
	return sf
}
