package common

import "os"

var initFunctions = make([]func(), 0)

func InitCommon() {
	for _, apply := range initFunctions {
		apply()
	}
}

func AddInitialized(apply func()) {
	initFunctions = append(initFunctions, apply)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}
