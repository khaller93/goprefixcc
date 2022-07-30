package main

import (
	"flag"
	"fmt"
	"github.com/khaller93/goprefixcc/api"
	"os"
)

const VERSION string = "1.1.0"

func printUsage(errorMessage string) {
	fmt.Printf("error: %s.\n\n", errorMessage)
	appName := "goprefixcc"
	if len(os.Args) > 1 {
		appName = os.Args[0]
	}
	fmt.Printf("Usage: %s (string [-reverse]) | -version\n", appName)
	flag.PrintDefaults()
}

func main() {
	version := flag.Bool("version", false, "prints version of this app")
	reverse := flag.Bool("reverse", false, "performs a reverse lookup")
	flag.Parse()

	if *version {
		fmt.Printf("goprefixcc v%v\n", VERSION)
		return
	}

	var args = flag.Args()
	if len(args) == 1 {
		var prefixCCAPI = api.GetOnTheFlyPrefixCC()
		// do forward or reverse lookup
		if *reverse {
			prefixListOpt, err := prefixCCAPI.GetPrefixName(args[0])
			if err == nil {
				if prefixListOpt.Found() {
					prefixList := prefixListOpt.Value()
					for i := 0; i < len(prefixList); i++ {
						fmt.Println(prefixList[i])
					}
				} else {
					fmt.Printf("no prefix could be found for '%s'\n", args[0])
					os.Exit(1)
				}
			} else {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		} else {
			namespaceListOpt, err := prefixCCAPI.GetNamespace(args[0])
			if err == nil {
				if namespaceListOpt.Found() {
					namespaceList := namespaceListOpt.Value()
					for i := 0; i < len(namespaceList); i++ {
						fmt.Println(namespaceList[i])
					}
				} else {
					fmt.Printf("no namespaces could be found for '%s'\n", args[0])
					os.Exit(1)
				}
			} else {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}
	} else {
		printUsage("exactly one string argument has to be passed to the application")
		os.Exit(1)
	}
}
