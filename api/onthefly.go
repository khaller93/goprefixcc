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
                    var lines = strings.Split(string(data[:]), "\n")
                    namespaces := make([]string, 0)
                    for i := 0; i < len(lines); i++ {
                        psplit := strings.Split(lines[i], "\t")
                        if len(psplit) == 2 && len(psplit[1]) > 0 {
                            namespaces = append(namespaces, psplit[1])
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
                    var lines = strings.Split(string(data[:]), "\n")
                    prefixList := make([]string, 0)
                    for i := 0; i < len(lines); i++ {
                        psplit := strings.Split(lines[i], "\t")
                        if len(psplit) == 2 && len(psplit[0]) > 0 {
                            prefixList = append(prefixList, psplit[0])
                        }
                    }
                    return prefixList, nil
                }
            }
        }
    }
    return nil, reverseLookupError{iri: iri};
}
