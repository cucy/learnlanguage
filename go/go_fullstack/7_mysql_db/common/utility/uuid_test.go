package utility

import (
	"testing"
	"fmt"
)

func TestGenerateUUID(t *testing.T) {
	for i := 0; i < 5; i++ {
		r_uuid := GenerateUUID()
		fmt.Println(r_uuid)
	}
}
