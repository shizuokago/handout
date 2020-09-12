package internal

import (
	"fmt"
	"net/url"
)

func Hello(p string, u *url.URL) []byte {
	return []byte(fmt.Sprintf("Hello %s[%s]", p, u.String()))
}
