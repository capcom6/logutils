# logutils

logutils is a Go package that augments the standard library "log" package
to make logging a bit more modern, without fragmenting the Go ecosystem
with new logging packages.

## The simplest thing that could possibly work

Presumably your application already uses the default `log` package. To switch, you'll want your code to look like the following:

```go
package main

import (
	"log"
	"os"

	"github.com/capcom6/logutils"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels: []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer: os.Stderr,
	}
	log.SetOutput(filter)

	log.Print("[DEBUG] Debugging") // this will not print
	log.Print("[WARN] Warning") // this will
	log.Print("[ERROR] Erring") // and so will this
	log.Print("Message I haven't updated") // and so will this
}
```

This logs to standard error exactly like go's standard logger. Any log messages you haven't converted to have a level will continue to print as before.

## Default filter

You can also use the default filter named `DefaultLevelFilter` as follows:

```go
package main

import (
	"log"

	"github.com/capcom6/logutils"
)

func main() {
	log.SetOutput(logutils.DefaultLevelFilter)

	logutils.Debug("Debugging") // this will not print
	logutils.Warn("Warning") // this will
	logutils.Error("Erring") // and so will this
	logutils.Print("Message I haven't updated") // and so will this
}

```
