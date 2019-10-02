package main

import (
	"flag"
	"fmt"
	"goprefixcc/api"
	"os"
)

const version string = "1.0"

func main() {
	version := flag.Bool("version", false, "prints version of this app")
	reverse := flag.Bool("reverse", false, "performs a reverse lookup")
	flag.Parse()

	if *version {
		fmt.Printf("(2019) goprefixcc version %v\n", version)
		return
	}

	var args []string = flag.Args()
	if len(args) == 1 {
		var prefixCCapi api.PrefixCC = api.GetOnTheFlyPrefixCC()
		// do forward or reverse lookup
		if *reverse {
			prefixList, err := prefixCCapi.GetPrefixName(args[0])
			if err == nil {
				for i := 0; i < len(prefixList); i++ {
					fmt.Println(prefixList[i])
				}
			} else {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			return
		} else {
			namespaceList, err := prefixCCapi.GetNamespace(args[0])
			if err == nil {
				for i := 0; i < len(namespaceList); i++ {
					fmt.Println(namespaceList[i])
				}
			} else {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			return
		}
	} else {
		fmt.Fprintln(os.Stderr, "Exactly one string argument has to be passed to the application.")
	}
}
