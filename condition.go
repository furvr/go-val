package val

// Condition defines a data + rule pair to test.
type Condition struct {
	Data *Data
	Rule ValidatorFunc
}

// NewCondition creates a new instance of a Condition.
func NewCondition(data *Data, rule ValidatorFunc) *Condition {
	return &Condition{
		Data: data,
		Rule: rule,
	}
}

// Validate returns a (potential) error value returned by evaluating the called
// conditions rule with the conditions data value.
func (c *Condition) Validate() error {
	return c.Rule(c.Data)
}
