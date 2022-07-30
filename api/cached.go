package api

import "github.com/khaller93/goprefixcc/utils"

// cached fetches namespaces and prefix from a data dump of the online service.
type cached struct {
	db *CachedDB
}

// GetCachedPrefixCC gets a PrefixCC instance that is using a cached
// version of the prefix.cc service.
func GetCachedPrefixCC() PrefixCC {
	return &cached{
		db: &prefixDump,
	}
}

func (c *cached) GetNamespace(prefix string) (*OptionalV, error) {
	val, found := c.db.Entries[prefix]
	if found {
		return OptionalOf([]string{val}), nil
	}
	return EmptyV, nil
}

func (c *cached) GetPrefixName(iri string) (*OptionalV, error) {
	namespace, err := utils.ExtractNamespaceInformation(iri)
	if err != nil {
		return nil, err
	}
	values := make([]string, 0)
	for key, val := range c.db.Entries {
		if val == namespace {
			values = append(values, key)
		}
	}
	if len(values) == 0 {
		return EmptyV, nil
	}
	return OptionalOf(values), nil
}
