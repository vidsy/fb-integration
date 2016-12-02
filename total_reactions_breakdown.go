package fbintegration

type (
	// TotalReactionsBreakdown slice of ReactionsBreakdowns
	TotalReactionsBreakdown []ReactionsBreakdown
)

// Swap interface method for sorting, swaps the items positions
// in the slice.
func (p TotalReactionsBreakdown) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Len interface method for sorting, gives length of items in slice.
func (p TotalReactionsBreakdown) Len() int {
	return len(p)
}

// Less interface method for sorting, finds which is lower out of the
// two given items.
func (p TotalReactionsBreakdown) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

// HasType checks if the map has a key for a given type.
func (p TotalReactionsBreakdown) HasType(typeName string) bool {
	for _, reactionBreakdown := range p {
		if reactionBreakdown.Type == typeName {
			return true
		}
	}

	return false
}

// IncrementValueForType increments the value for a given type, such as 'WOW' or 'LIKE'.
func (p TotalReactionsBreakdown) IncrementValueForType(typeName string, amount float64) {
	for i, reactionBreakdown := range p {
		if typeName == reactionBreakdown.Type {
			p[i].Increment(amount)
		}
	}
}
