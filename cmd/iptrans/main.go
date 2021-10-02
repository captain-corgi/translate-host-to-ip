package main

import (
	"fmt"
	"os"

	"github.com/captain-corgi/translate-host-to-ip/pkg/iptrans"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("URL missing. Usage: \nmake url=<your url>")
	} else {
		fmt.Println(iptrans.Lookup(argsWithoutProg[0]))
	}
}
