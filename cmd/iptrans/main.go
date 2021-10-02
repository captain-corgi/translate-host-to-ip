package main

import (
	"fmt"

	"github.com/captain-corgi/translate-host-to-ip/pkg/iptrans"
)

func main() {
	fmt.Println(iptrans.Lookup("1www.google.com"))
}
