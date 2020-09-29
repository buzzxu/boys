package bcrypts

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	salt, _ := Salt()

	salt, _ = Salt(10)

	password := "123456"
	hash, _ := Hash(password)
	if Match(password, hash) {
		fmt.Println("They match")
	}

	hash, _ = Hash(password, salt)
	if Match(password, hash) {
		fmt.Println("They match")
	}

	bad_password := "123455"
	if !Match(bad_password, hash) {
		fmt.Println("They don't match")
	}
	fmt.Printf("%t \n", Match("111111", "$2a$15$u7OC5fL8xNQQSWP7MoJCyerLqx.H73G1JFUuzw6ztRELiwAGDVKtu"))
}
