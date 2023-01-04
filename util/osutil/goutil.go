package osutil

import (
	"ftools/util/cmdutil"
	"github.com/pkg/errors"
	"regexp"
	"runtime"
	"strings"
)

// GoVersion 获取go版本。例如：“1.18.2”
func GoVersion() string {
	return runtime.Version()[2:]
}

// GoInfo define
type GoInfo struct {
	Version string
	GoOS    string
	Arch    string
}

// match "go version go1.19 darwin/amd64"
var goVerRegex = regexp.MustCompile(`\sgo([\d.]+)\s(\w+)/(\w+)`)

// ParseGoVersion 通过解析“go version”结果获取信息。
func ParseGoVersion(line string) (*GoInfo, error) {
	// eg: [" go1.19 darwin/amd64", "1.19", "darwin", "amd64"]
	lines := goVerRegex.FindStringSubmatch(line)
	if len(lines) != 4 {
		return nil, errors.New("returns go info is not full")
	}

	info := &GoInfo{}
	info.Version = strings.TrimPrefix(lines[1], "go")
	info.GoOS = lines[2]
	info.Arch = lines[3]

	return info, nil
}

// OsGoInfo 获取并解析
func OsGoInfo() (*GoInfo, error) {
	cmdArgs := []string{"env", "GOVERSION", "GOOS", "GOARCH"}
	line, err := cmdutil.ExecCmd("go", cmdArgs)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(line), "\n")

	if len(lines) != len(cmdArgs)-1 {
		return nil, errors.New("returns go info is not full")
	}

	info := &GoInfo{}
	info.Version = strings.TrimPrefix(lines[0], "go")
	info.GoOS = lines[1]
	info.Arch = lines[2]

	return info, nil
}
