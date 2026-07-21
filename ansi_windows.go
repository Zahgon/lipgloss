//go:build windows

package lipgloss

import (
	"os"
)

func EnableLegacyWindowsANSI(f *os.File) { _ = "STUB: not implemented"; return }
