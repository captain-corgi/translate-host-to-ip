# TRANSLATE HOST TO IP

<a href='https://github.com/captain-corgi/translate-host-to-ip' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-100%25-brightgreen.svg?longCache=true&style=flat)</a>

This is a small tool for translate host domain into IP address.

Usage:

    make url="your url"

Example:
    
    make url=google.com
    result: 74.125.24.113

## Import into your project:
Import:
> go get github.com/captain-corgi/translate-host-to-ip

Usage:
<pre>
import "github.com/captain-corgi/translate-host-to-ip/pkg/iptrans"
....
iptrans.Lookup("url")
</pre>
