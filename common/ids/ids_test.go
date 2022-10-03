package ids

import (
	"os"
	"testing"
)

func TestSnowflake_Generate(t *testing.T) {
	snowflake := SnowflakeID{}
	os.Setenv("IDS_NODE_NO", "1")
	snowflake.Init()
	println(snowflake.Generate())
}
