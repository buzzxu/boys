package ids

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
	"strconv"
	"sync"
)

type SnowflakeID struct {
	node *snowflake.Node
	once sync.Once
}

func (s *SnowflakeID) Init() error {
	key, ok := os.LookupEnv("IDS_NODE_NO")
	if !ok {
		return fmt.Errorf("IDS_NODE_NO is not set in system environment")
	}
	nodeNo, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return err
	}
	s.New(nodeNo)
	return nil
}
func (s *SnowflakeID) New(node int64) {
	s.once.Do(func() {
		n, err := snowflake.NewNode(node)
		if err != nil {
			fmt.Println(err)
			return
		}
		s.node = n
	})
}

func (s *SnowflakeID) Generate() string {
	if s.node == nil {
		panic("Snowflake node is nil")
		return ""
	}
	return s.node.Generate().String()
}
