# simplelogger
Really simple logger inspired by Dave Cheney article at http://dave.cheney.net/2015/11/05/lets-talk-about-logging.

## Usage
Here is a sample adapted from the test code:
```
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   true,
		DEBUG:  false,
		Writer:   buf,
	}

	sl.Info("[info] %s\n", "foo")
	sl.Debug("[debug] %s\n", "bar")
	expected := "[info] foo\n"
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %#v, got %#v", expected, result)
	}
```

You first decide and create the io.Writer that determines where
the log output should go. Then you allocate the SimpleLogger struct using the literal syntax, setting INFO and DEBUG to true as desired.

