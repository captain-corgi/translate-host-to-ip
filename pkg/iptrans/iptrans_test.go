package iptrans

import (
	"fmt"
	"net"
	"testing"
)

var mockedLookupIP = func(host string) ([]net.IP, error) {
	switch host {
	case "localhost":
		return []net.IP{net.IPv4(127, 0, 0, 1)}, nil
	case "google.com":
		return []net.IP{net.IPv4(127, 0, 0, 1)}, nil
	default:
		return []net.IP{}, fmt.Errorf("Some thing wrong")
	}
}

func TestLookup(t *testing.T) {
	// Init Mock function
	netLIP = mockedLookupIP

	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1. Success with localhost",
			args: args{
				domain: "localhost",
			},
			want: "127.0.0.1",
		},
		{
			name: "2. Success with google",
			args: args{
				domain: "google.com",
			},
			want: "127.0.0.1",
		},
		{
			name: "3. Domain not correct",
			args: args{
				domain: "aaagoogle.com",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lookup(tt.args.domain); got != tt.want {
				t.Errorf("Lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
