package val

// Contract represents an agreement between
type Contract struct {
	terms []*Condition
	Fails []*Failure
}

// NewContract DOC: ..
func NewContract() *Contract {
	return &Contract{}
}

// AddRule DOC: ..
func (c *Contract) AddRule(data *Data, rule ValidatorFunc) {
	c.terms = append(c.terms, NewCondition(data, rule))
}

// Validate iterates all terms, calling each condition's `Validate` function,
// and appends a `Failure` to `c.Fails` when a condition returns an error. The
// returned boolean will be false when one or more terms are not met.
func (c *Contract) Validate() bool {
	var fcount = 0

	for _, cond := range c.terms {
		if err := cond.Validate(); err != nil {
			c.Fails = append(c.Fails, NewFailure(cond, err))
			fcount = fcount + 1
		}
	}

	if fcount > 0 {
		return false
	}

	return true
}
