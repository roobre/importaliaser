package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"roob.re/importaliaser"
)

func main() {
	jsonPath := flag.String("j", "", "Path to json store")
	addr := flag.String("a", "", "Address to listen on")
	flag.Parse()

	if *jsonPath == "" || *addr == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s -a :8080 -j /path/to/json\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	aliaser := importaliaser.NewAliaser(importaliaser.NewJSONStorer(*jsonPath))
	http.ListenAndServe(*addr, aliaser)
}
