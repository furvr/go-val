package val

import "testing"

// ---

// TestEqualsString tests the EqualsString rule on a fresh contract.
func TestEqualsString(t *testing.T) {
	var contract = NewContract()
	var value = "value"
	var data = NewData([]byte(value), []string{"TEST-VALUE"})

	contract.AddRule(data, EqualsString(value, "VALUE-NOT-EQUALS"))

	if pass := contract.Validate(); !pass {
		for _, f := range contract.Fails {
			t.Errorf("Expected rules to pass; Data with sig `%v` returned errors: %v",
				f.Condition.Data.Signature,
				f.Error)
		}
	}
}
