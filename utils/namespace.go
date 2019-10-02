package utils

import "regexp"

var prefixPattern, prefixPatternError = regexp.Compile("^(.*[:#/])(.*[^:#/])?$")

type namespaceExtractionError struct {
	iri string
}

func (e namespaceExtractionError) Error() string {
	return "Could not extract utils from the IRI '" + e.iri + "'."
}

// Tries to extract utils from a given IRI. This
// method returns the utils IRI, if it could be
// extracted or an error otherwise.
func ExtractNamespaceInformation(iri string) (string, error) {
	if prefixPatternError == nil && prefixPattern != nil {
		subMatches := prefixPattern.FindAllStringSubmatch(iri, 1)
		if len(subMatches) > 0 {
			return subMatches[0][1], nil
		}
	}
	return "", namespaceExtractionError{iri: iri}
}
