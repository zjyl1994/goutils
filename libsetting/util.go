package libsetting

import (
	"errors"
	"os"
	"strings"
)

func fileExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return !errors.Is(err, os.ErrNotExist)
	}
}

func searchConf() (filename string, ok bool) {
	appName := strings.ToLower(applicationName)
	// search in env var
	if filename, ok = envData["setting"]; ok {
		return filename, true
	}
	// search in etc
	filename = "/etc/" + appName + "/" + appName + ".conf"
	if fileExist(filename) {
		return filename, true
	}
	filename = "/etc/" + appName + "/config.conf"
	if fileExist(filename) {
		return filename, true
	}
	// search in workdir
	filename = appName + ".conf"
	if fileExist(filename) {
		return filename, true
	}
	return "", false
}

func cutWithStr(srcStr, cutStr string) (string, string) {
	equalMarkPos := strings.Index(srcStr, cutStr)
	if equalMarkPos > 0 {
		return srcStr[:equalMarkPos], srcStr[equalMarkPos+len(cutStr):]
	} else {
		return srcStr, ""
	}
}
