package api

import (
	"github.com/hallerkevin/goprefixcc/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// prefix CC type that fetches namespaces and
// prefix on the fly.
type onTheFly struct {
	url string
}

// gets the prefix CC type that fetches namespaces and
// prefix on the fly.
func GetOnTheFlyPrefixCC() PrefixCC {
	return onTheFly{url: PrefixCCAPIURL}
}

func (o onTheFly) GetNamespace(prefix string) ([]string, error) {
	req, err := http.NewRequest("GET", o.url+prefix+".file.txt", nil)
	if err == nil && req != nil {
		hc := http.Client{}
		resp, err := hc.Do(req)
		if err == nil && resp != nil {
			var status int = resp.StatusCode
			if status >= 200 && status < 300 {
				data, err := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()
				if err == nil {
					var lines []string = strings.Split(string(data[:]), "\n")
					namespaces := make([]string, len(lines))
					for i := 0; i < len(lines); i++ {
						psplit := strings.Split(lines[i], "\t")
						if len(psplit) == 2 {
							namespaces[i] = psplit[1]
						}
					}
					return namespaces, nil
				}
			}
		}
	}
	return nil, namespaceLookupError{prefix: prefix}
}

func (o onTheFly) GetPrefixName(iri string) ([]string, error) {
	namespace, err := utils.ExtractNamespaceInformation(iri)
	if err == nil {
		resp, err := http.Get(o.url + "reverse?format=txt&uri=" + url.QueryEscape(namespace))
		if err == nil && resp != nil {
			var status int = resp.StatusCode
			if status >= 200 && status < 300 {
				data, err := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()
				if err == nil {
					var lines []string = strings.Split(string(data[:]), "\n")
					prefixList := make([]string, len(lines))
					for i := 0; i < len(lines); i++ {
						psplit := strings.Split(lines[i], "\t")
						if len(psplit) == 2 {
							prefixList[i] = psplit[0]
						}
					}
					return prefixList, nil
				}
			}
		}
	}
	return nil, reverseLookupError{iri: iri};
}
