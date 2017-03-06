package val

// Failure is a representation of a failed condition and it's returned errors.
type Failure struct {
	Condition *Condition
	Error     error
}

// NewFailure creates an instance of a Failure.
func NewFailure(cond *Condition, errs error) *Failure {
	return &Failure{
		Condition: cond,
		Error:     errs,
	}
}
