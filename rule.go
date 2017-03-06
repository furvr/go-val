package val

import "errors"

// ---

// ValidatorFunc defines a function signature for Validation Rules
type ValidatorFunc func(*Data) error

// ---

// EqualsString takes a value string and a failure message string. The returned
// ValidatorFunc tests for string equality, returning the provided failmsg as a
// standard Go error when the equals comparison fails.
func EqualsString(val string, failmsg string) ValidatorFunc {
	return func(data *Data) error {
		if data.String() != val {
			return errors.New(failmsg)
		}

		return nil
	}
}
