package api

import (
	"github.com/khaller93/goprefixcc/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const PrefixCCAPIURL string = "https://prefix.at/"

// onTheFly fetches namespaces and prefix from the online service.
type onTheFly struct {
	url     string
	timeout time.Duration
}

// GetOnTheFlyPrefixCC gets a PrefixCC instance that fetches namespaces and
// prefix from the online service.
func GetOnTheFlyPrefixCC() PrefixCC {
	return &onTheFly{
		url:     PrefixCCAPIURL,
		timeout: 10 * time.Second,
	}
}

func (o *onTheFly) GetNamespace(prefix string) (*OptionalV, error) {
	req, err := http.NewRequest("GET", o.url+prefix+".file.txt", nil)
	if err == nil && req != nil {
		hc := http.Client{}
		resp, err := hc.Do(req)
		if err == nil && resp != nil {
			var status = resp.StatusCode
			if status >= 200 && status < 300 {
				data, err := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()
				if err == nil {
					var lines = strings.Split(string(data[:]), "\n")
					namespaces := make([]string, 0)
					for i := 0; i < len(lines); i++ {
						pSplit := strings.Split(lines[i], "\t")
						if len(pSplit) == 2 && len(pSplit[1]) > 0 {
							namespaces = append(namespaces, pSplit[1])
						}
					}
					return OptionalOf(namespaces), nil
				}
				return nil, err
			}
			return EmptyV, nil
		}
		return nil, err
	}
	return nil, err
}

func (o *onTheFly) GetPrefixName(iri string) (*OptionalV, error) {
	namespace, err := utils.ExtractNamespaceInformation(iri)
	if err == nil {
		resp, err := http.Get(o.url + "reverse?format=txt&uri=" + url.QueryEscape(namespace))
		if err == nil && resp != nil {
			var status = resp.StatusCode
			if status >= 200 && status < 300 {
				data, err := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()
				if err == nil {
					var lines = strings.Split(string(data[:]), "\n")
					prefixList := make([]string, 0)
					for i := 0; i < len(lines); i++ {
						pSplit := strings.Split(lines[i], "\t")
						if len(pSplit) == 2 && len(pSplit[0]) > 0 {
							prefixList = append(prefixList, pSplit[0])
						}
					}
					return OptionalOf(prefixList), nil
				}
				return nil, err
			}
			return EmptyV, nil
		}
		return nil, err
	}
	return nil, err
}
