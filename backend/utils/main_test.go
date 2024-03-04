package util

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// TestMain: the main test driver.
//
// returns: nothing
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
