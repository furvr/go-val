package val

// Data is a representation of testable data. Testable data includes a value of
// bytes and an arbitrary signature containing strings which should identify the
// data value or its characteristics outside of the local scope.
type Data struct {
	Value     []byte
	Signature []string
}

// NewData returns an instance of `Data` with a value and a signature.
func NewData(val []byte, sig []string) *Data {
	return &Data{
		Value:     val,
		Signature: sig,
	}
}

// String returns a string from the current data value.
func (d *Data) String() string {
	return string(d.Value)
}
