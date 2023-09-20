package logutils

import "os"

var DefaultLevelFilter = &LevelFilter{
	Levels:   []LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"},
	MinLevel: "WARN",
	Writer:   os.Stdout,
}
