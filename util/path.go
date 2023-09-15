package util

import (
	"path/filepath"
	"runtime"
)

func GetProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return filepath.Join(basepath, "..")
}
