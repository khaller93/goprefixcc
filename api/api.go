package api

// PrefixCC is an API to the prefix.cc service.
type PrefixCC interface {
	// GetNamespace aims to get the namespace IRI for the given prefix. If
	// at least one namespace can be found, then a non-empty list of suggested
	// IRIs wrapped in an OptionalV and nil will be returned. EmptyV and nil will
	// be returned, if the prefix is unknown. Otherwise, if accessing the service
	// fails, then nil and the corresponding error will be returned.
	GetNamespace(prefix string) (*OptionalV, error)
	// GetPrefixName aims to get the prefix for the given IRI. If at least one
	// prefix can be found, then a non-empty list of suggested prefixes wrapped
	// in an OptionalV and nil will be returned. Otherwise, if accessing the service
	// fails, then nil and the corresponding error will be returned.
	GetPrefixName(iri string) (*OptionalV, error)
}

// OptionalV is an object that wraps a value, or represents
// the state of "no value could be found".
type OptionalV struct {
	values []string
}

// Value returns a non-empty array of strings, if Found() method returns
// true. Otherwise, nil will be returned.
func (opt OptionalV) Value() []string {
	return opt.values
}

// Found returns true, if a value can be received by calling the
// Value() method, otherwise false.
func (opt OptionalV) Found() bool {
	return opt.values != nil
}

// OptionalOf returns a OptionalV with the given values wrapped.
func OptionalOf(values []string) *OptionalV {
	return &OptionalV{values: values}
}

// EmptyV is an OptionalV, which represents the state of "no value could be found".
var EmptyV = &OptionalV{
	values: nil,
}
