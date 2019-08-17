package main

import (
	"flag"
	"fmt"
	"github.com/tcnksm/go-latest"
)

var version = "1.0.0"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}

	githubTag := &latest.GithubTag{
		Owner: "raimon49",
		Repository: "go-for-everyone-reading",
	}

	res, _ := latest.Check(githubTag, "v20190526")
	if res.Outdated {
		fmt.Printf("%s is not latest, you should upgrade to %s", version, res.Current)
	}
}
