package ids

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSnowflake_Generate(t *testing.T) {
	snowflake := SnowflakeID{}
	os.Setenv("IDS_NODE_NO", "1")
	snowflake.Init()
	println(snowflake.Generate())
}

func TestIdWorker_NextId(t *testing.T) {
	worker, err := NewIDWorker()
	if err != nil {
		log.Fatalf("Failed to create IdWorker: %v", err)
	}

	for i := 0; i < 10; i++ {
		id, err := worker.NextID()
		if err != nil {
			log.Fatalf("Failed to generate id: %v", err)
		}
		fmt.Println(id)
	}
}
