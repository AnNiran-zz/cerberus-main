package config

import "os"

var (
	mainPath = os.Getenv("GOPATH") + "/src/cerberus/"
	LogsPath = mainPath + "/log/"
)
