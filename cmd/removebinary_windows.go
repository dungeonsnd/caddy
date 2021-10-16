//go:build windows
// +build windows

package caddycmd

import (
	"os"
	"path/filepath"
	"syscall"
)

// removeCaddyBinary removes the Caddy binary at the given path.
//
// On Windows, this uses a syscall to indirectly remove the file,
// because otherwise we get an "Access is denied." error when trying
// to delete the binary while Caddy is still running and performing
// the upgrade. "cmd.exe /C" executes a command specified by the
// following arguments, i.e. "del" which will run as a separate process,
// which avoids the "Access is denied." error.
func removeCaddyBinary(path string) error {
	var sI syscall.StartupInfo
	var pI syscall.ProcessInformation
	argv := syscall.StringToUTF16Ptr(filepath.Join(os.Getenv("windir"), "system32", "cmd.exe") + " /C del " + path)
	return syscall.CreateProcess(nil, argv, nil, nil, true, 0, nil, nil, &sI, &pI)
}
