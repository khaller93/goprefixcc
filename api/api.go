package api

const PrefixCCAPIURL string = "https://prefix.cc/"

// API type for fetching prefix.cc data
type PrefixCC interface {
    GetNamespace(prefix string) ([]string, error)
    GetPrefixName(iri string) ([]string, error)
}

// error when fetching namespaces for a prefix
type namespaceLookupError struct {
    prefix string
}

func (e namespaceLookupError) Error() string {
    return "Could not fetch namespaces for prefix '" + e.prefix + "'."
}

// error when doing a reverse lookup
type reverseLookupError struct {
    iri string
}

func (e reverseLookupError) Error() string {
    return "Could not fetch prefix names for namespace '" + e.iri + "'."
}

