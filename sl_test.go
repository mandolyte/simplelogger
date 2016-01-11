package simplelogger

import (
	"bytes"
	"testing"
)

func TestInfoMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   true,
		DEBUG:  false,
		Writer:   buf,
	}

	sl.Info("foo")
	sl.Debug("bar")

	expected := "foo"
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %#v", expected, result)
	}
}

func TestFormattedInfoMessages(t *testing.T) {
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
}

func TestDebugMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   false,
		DEBUG:  true,
		Writer:   buf,
	}

	sl.Info("foo")
	sl.Debug("bar")

	expected := "bar"
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %s", expected, result)
	}
}

func TestFormattedDebugMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   false,
		DEBUG:  true,
		Writer:   buf,
	}

	sl.Info("[info] %s\n", "foo")
	sl.Debug("[debug] %s\n", "bar")
	expected := "[debug] bar\n"
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %s", expected, result)
	}
}

func TestBothMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   true,
		DEBUG:  true,
		Writer:   buf,
	}

	sl.Info("foo")
	sl.Debug("bar")

	expected := "foobar"
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %s", expected, result)
	}
}

func TestFormattedBothMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   true,
		DEBUG:  true,
		Writer:   buf,
	}

	sl.Info("[info] %s\n", "foo")
	sl.Debug("[debug] %s\n", "bar")
	expected := "[info] foo\n[debug] bar\n"
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %s", expected, result)
	}
}

func TestNoMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   false,
		DEBUG:  false,
		Writer:   buf,
	}

	sl.Info("foo")
	sl.Debug("bar")

	expected := ""
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %s", expected, result)
	}
}

func TestFormattedNoMessages(t *testing.T) {
	buf := new(bytes.Buffer)
	sl := &SimpleLogger{
		INFO:   false,
		DEBUG:  false,
		Writer:   buf,
	}

	sl.Info("[info] %s\n", "foo")
	sl.Debug("[debug] %s\n", "bar")
	expected := ""
	result := buf.String()
	if result != expected {
		t.Fatalf("[fail] expected: %s, got %s", expected, result)
	}
}

