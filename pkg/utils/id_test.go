package utils_test

import (
	"fmt"
	"github.com/lovelacelee/clsgo/pkg/utils"
	"testing"
)

func TestSessionId(t *testing.T) {
	for i := 0; i < 2; i++ {
		fmt.Println(utils.SessionId())
		fmt.Println(utils.SessionId(8))
	}
	fmt.Println(utils.UUID())
	fmt.Println(utils.UUIDV1())
}
