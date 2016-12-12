package fbintegration_test

import (
	"github.com/vidsy/fbintegration"
	"testing"
)

func TestReactionsBreakdown(t *testing.T) {
	t.Run(".Increment()", func(t *testing.T) {
		t.Run("IncrementsValue", func(t *testing.T) {
			reactionsBreakdown := fbintegration.ReactionsBreakdown{"LIKE", 1}
			reactionsBreakdown.Increment(19)

			if reactionsBreakdown.Value != 20 {
				t.Fatalf("Expected .Value to be 20, got: %.2f", reactionsBreakdown.Value)
			}
		})
	})
}
