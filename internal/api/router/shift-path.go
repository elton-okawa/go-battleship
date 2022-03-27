package router

import (
	"path"
	"strings"
)

// Splits given path into <head>/<tail>
// Example - /users
// - /users -> users, /
// - / -> "", /
// Example - /users/10
// - /users/10 -> users, /10
// - /10 -> 10, /
// Example - /users/10/receipts
// - /users/10/receipts -> users, /10/receipts
// - /10/receipts -> 10, /receipts
// - /receipts -> receipts, /
//
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
