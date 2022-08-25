package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equal(t *testing.T, actual any, wanted any) {
	if !cmp.Equal(actual, wanted) {
		t.Errorf("comparison failed: %s", cmp.Diff(actual, wanted))
	}
}
