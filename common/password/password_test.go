package password

import "testing"

func TestCheckPassword(t *testing.T) {
	println(CheckPassword("111111", "$2a$10$sNqAI1td8P5Kroc4SzpiaewjFYQk81Dsj.fc9C3MvSdb1WOFypPxm"))
}
