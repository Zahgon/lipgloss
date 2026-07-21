//go:build !windows

package lipgloss

import "os"

func EnableLegacyWindowsANSI(*os.File) { _ = "STUB: not implemented"; return }
