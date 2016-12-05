package fbintegration_test

import (
	"testing"

	"github.com/vidsy/fbintegration"
	facebookLib "github.com/huandu/facebook"
)

func TestPostsSummary(t *testing.T) {
	t.Run("NewPostsSummary()", func(t *testing.T) {
		t.Run("CorrectTotalReactionsBreandown", func(t *testing.T) {
			var results []*facebookLib.Result
			result := facebookLib.MakeResult(byte[](`{"data" : [],"summary" : {"total_count" : 2}}`)) 


			//results := PostResults.ReactionBreakdown = 
		})
	})
}
