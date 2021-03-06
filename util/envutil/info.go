package envutil

import (
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/preakrel/go-config/util/sysutil"
)

// IsWin system. linux windows darwin
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac system
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux system
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsConsole 判断 w 是否为 stderr、stdout、stdin 三者之一
func IsConsole(out io.Writer) bool {
	o, ok := out.(*os.File)
	if !ok {
		return false
	}

	return o == os.Stdout || o == os.Stderr || o == os.Stdin
}

// IsMSys msys(MINGW64) 环境，不一定支持颜色
func IsMSys() bool {
	// "MSYSTEM=MINGW64"
	if len(os.Getenv("MSYSTEM")) > 0 {
		return true
	}

	return false
}

// HasShellEnv has shell env check.
// Usage:
// 	HasShellEnv("sh")
// 	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	return sysutil.HasShellEnv(shell)
}

// IsSupportColor check console is support color.
// supported: linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// not support: windows cmd, powerShell
func IsSupportColor() bool {
	// "TERM=xterm"  support color
	// "TERM=xterm-vt220" support color
	// "TERM=xterm-256color" support color
	// "TERM=cygwin" don't support color
	if strings.Contains(os.Getenv("TERM"), "xterm") {
		return true
	}

	// like on ConEmu software, e.g "ConEmuANSI=ON"
	if os.Getenv("ConEmuANSI") == "NO" {
		return true
	}

	// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"
	if os.Getenv("ANSICON") != "" {
		return true
	}

	return false
}

// IsSupport256Color is support 256 color
func IsSupport256Color() bool {
	// "TERM=xterm-256color"
	return strings.Contains(os.Getenv("TERM"), "256color")
}
