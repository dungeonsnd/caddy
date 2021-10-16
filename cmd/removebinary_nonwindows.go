//go:build !windows
// +build !windows

package caddycmd

import (
	"os"
)

// removeCaddyBinary removes the Caddy binary at the given path.
//
// On any non-Windows OS, this simply calls os.Remove, since they should
// probably not exhibit any issue with processes deleting themselves.
func removeCaddyBinary(path string) error {
	return os.Remove(path)
}
