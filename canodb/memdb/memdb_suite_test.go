package memdb

import (
	"testing"

	"github.com/lindsuen/canodb/canodb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
