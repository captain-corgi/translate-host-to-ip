package iptrans

import (
	"fmt"
	"net"
)

func Lookup(domain string) string {
	ip, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return ""
	}
	return fmt.Sprint(ip[0])
}
