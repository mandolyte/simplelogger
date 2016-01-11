package main

import "os"
import "github.com/mandolyte/simplelogger"

func main() {
	sl := &simplelogger.SimpleLogger{
		INFO:   true,
		DEBUG:  false,
		Writer: os.Stderr,
	}

	sl.Info("[info] %s\n", "foo") // will print
	sl.Debug("[debug] %s\n", "bar") // will not print

	sl.DEBUG = true
	sl.Debug("[debug] %s\n", "barnone") // will now print
}
