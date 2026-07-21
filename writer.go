package lipgloss

import (
	"io"
	"os"

	"github.com/charmbracelet/colorprofile"
)

var Writer = colorprofile.NewWriter(os.Stdout, os.Environ())

func Println(v ...any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func Printf(format string, v ...any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func Print(v ...any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func Fprint(w io.Writer, v ...any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func Fprintln(w io.Writer, v ...any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func Fprintf(w io.Writer, format string, v ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:wrapcheck

func Sprint(v ...any) string { _ = "STUB: not implemented"; return "" }

//nolint:errcheck

func Sprintln(v ...any) string { _ = "STUB: not implemented"; return "" }

//nolint:errcheck

func Sprintf(format string, v ...any) string { _ = "STUB: not implemented"; return "" }

//nolint:errcheck
