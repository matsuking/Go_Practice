

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)

	// テストをしたい関数を書いていく
	// main()
}
