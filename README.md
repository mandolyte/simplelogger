# simplelogger
Really simple logger inspired by Dave Cheney article at http://dave.cheney.net/2015/11/05/lets-talk-about-logging.

## Usage
You first decide and create the io.Writer that determines where
the log output should go. Then you allocate the SimpleLogger struct using the literal syntax, setting INFO and DEBUG to true as desired.

Here is an example from the "cmd" folder:
```
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

```
Here is sample output:
```
$ go run main.go
[info] foo
[debug] barnone
$ 
```