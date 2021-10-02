package iptrans

import (
	"fmt"
	"net"
)

var netLIP = net.LookupIP

func Lookup(domain string) string {
	ip, err := netLIP(domain)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return ""
	}
	return fmt.Sprint(ip[0])
}
