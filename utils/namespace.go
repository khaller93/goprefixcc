package utils

import "regexp"

var prefixPattern = regexp.MustCompile("^(.*[:#/])(.*[^:#/])?$")

type namespaceExtractionError struct {
	iri string
}

func (e namespaceExtractionError) Error() string {
	return "could not extract utils from the IRI '" + e.iri + "'"
}

// ExtractNamespaceInformation tries to extract the namespace from a given IRI.
// It returns the namespace IRI, if it could be extracted or an error otherwise.
func ExtractNamespaceInformation(iri string) (string, error) {
	subMatches := prefixPattern.FindAllStringSubmatch(iri, 1)
	if len(subMatches) > 0 {
		return subMatches[0][1], nil
	}
	return "", namespaceExtractionError{iri: iri}
}
