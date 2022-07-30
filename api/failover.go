package api

// failOver will use a onTheFly instance first, and if it fails, then
// the cached instance is used as a backup.
type failOver struct {
	onTheFly PrefixCC
	cached   PrefixCC
}

// GetFailOverPrefixCC will return a PrefixCC instance, which firstly uses an
// instance that makes use of the prefix.cc online service. If it fails, then
// an instance making use of a data dump from the service will be used.
func GetFailOverPrefixCC() PrefixCC {
	return &failOver{
		onTheFly: GetOnTheFlyPrefixCC(),
		cached:   GetCachedPrefixCC(),
	}
}

func (f *failOver) GetNamespace(prefix string) (*OptionalV, error) {
	val, err := f.onTheFly.GetNamespace(prefix)
	if err != nil {
		val, err = f.cached.GetNamespace(prefix)
		return val, err
	}
	return val, nil
}

func (f *failOver) GetPrefixName(iri string) (*OptionalV, error) {
	val, err := f.onTheFly.GetPrefixName(iri)
	if err != nil {
		val, err = f.cached.GetPrefixName(iri)
		return val, err
	}
	return val, nil
}
