# Val
Val is a small, flexible data validation library written in Go.

## Working with Contracts
A contract, in val-speak, is a list of terms (`[]Conditions`) which must be met. To create a contract, do:

```go
var contract = val.NewContract()
```

Add rules to a contract with `contract.AddRule(*val.Data, val.ValidatorFunc)`. In the following example `val.EqualsString` is a built-in validation function which validates that the value contained in the provided data equals the string provided in its first parameter.

```go
var data = NewData([]byte("value"), "DATA-TYPE", "DATA-CATEGORY")
contract.AddRule(data, val.EqualsString("value", "Values are not equal."))
```

To run validation, call `contract.Validate()`. The returned boolean indicates whether the contract passed all rules (`true`) or failed one or more rule (`false`).

When a contract fails validation, pointers to each failed `Condition` are provided in `contract.Fails`.

```go
if pass := contract.Validate(); !pass {
	for _, f := range contract.Fails {
		fmt.Printf("Expected rules to pass; Data with sig `%v` returned errors: %v",
			f.Condition.Data.Signature,
			f.Error)
	}
}
```

## Contributing

Pull-requests are welcome, and so are you. :)

Currently, approved PRs are those which improve performance and documentation, add useful functionality (core lib or built-ins), or fix bugs / flaws.

If you'd like to submit a PR which changes the API, adds currently unrepresented functionality, or modifies existing (stable) behavior, create an issue and reference your pull request so we can talk about it.

These contribution guidelines may be more structured in the future, depending on community reception.
