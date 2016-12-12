package fbintegration_test

import (
	"sort"
	"testing"

	"github.com/vidsy/fbintegration"
)

func TestTotalReactionsBreakdown(t *testing.T) {
	t.Run("Sort", func(t *testing.T) {
		t.Run("SortsByValue", func(t *testing.T) {
			var totalReactionsBreakdown fbintegration.TotalReactionsBreakdown
			totalReactionsBreakdown = append(totalReactionsBreakdown, fbintegration.ReactionsBreakdown{"WOW", 1})
			totalReactionsBreakdown = append(totalReactionsBreakdown, fbintegration.ReactionsBreakdown{"WOW", 5})

			sort.Sort(&totalReactionsBreakdown)

			if totalReactionsBreakdown[1].Value != 5.00 {
				t.Fatalf("Expected .Value to be 5, got: %.2f", totalReactionsBreakdown[1].Value)
			}
		})
	})

	t.Run(".HasType()", func(t *testing.T) {
		var totalReactionsBreakdown fbintegration.TotalReactionsBreakdown
		totalReactionsBreakdown = append(totalReactionsBreakdown, fbintegration.ReactionsBreakdown{"WOW", 5})

		t.Run("WithValidKey", func(t *testing.T) {
			if !totalReactionsBreakdown.HasType("WOW") {
				t.Fatalf(`Expected .HasType("WOW") to be true, got: %t`, totalReactionsBreakdown.HasType("WOW"))
			}
		})

		t.Run("WithMissingKey", func(t *testing.T) {
			if totalReactionsBreakdown.HasType("LIKE") {
				t.Fatalf(`Expected .HasType("LIKE") to be false, got: %t`, totalReactionsBreakdown.HasType("LIKE"))
			}
		})
	})

	t.Run(".IncrementValueForType()", func(t *testing.T) {
		var totalReactionsBreakdown fbintegration.TotalReactionsBreakdown
		totalReactionsBreakdown = append(totalReactionsBreakdown, fbintegration.ReactionsBreakdown{"WOW", 1})

		totalReactionsBreakdown.IncrementValueForType("WOW", 1)

		if totalReactionsBreakdown[0].Value != 2 {
			t.Fatalf(`Expected [0].Value to be 2, got: %.2f`, totalReactionsBreakdown[0].Value)
		}
	})
}
