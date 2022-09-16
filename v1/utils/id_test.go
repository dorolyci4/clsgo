package utils_test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
	"testing"
)

func TestSessionId(t *testing.T) {
	log.Green("Running id test cases")
	gtest.C(t, func(t *gtest.T) {
		id := utils.SessionId()
		for i := 0; i < 2; i++ {
			t.AssertNE(utils.SessionId(), id)
			t.Assert(len(utils.SessionId()), 44)
			t.Assert(len(utils.SessionId(8)), 8)
		}
		t.Assert(len(utils.UUID()), len("ce7197ab-4452-41d6-b803-0ecf24d0c0d1"))
		t.Assert(len(utils.UUIDV1()), len("ce7197ab-4452-41d6-b803-0ecf24d0c0d1"))
	})
}
