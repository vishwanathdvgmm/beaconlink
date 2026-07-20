package logging

import (
	"bytes"
	"strings"
	"testing"
)

func TestJSONLogger(t *testing.T) {

	var buf bytes.Buffer

	log, err := NewWithWriter(Config{
		Level:  "info",
		Format: "json",
	}, &buf)

	if err != nil {
		t.Fatal(err)
	}

	log.Info("hello")

	if !strings.Contains(buf.String(), "hello") {
		t.Fatal("log message missing")
	}
}

func TestTextLogger(t *testing.T) {

	var buf bytes.Buffer

	log, err := NewWithWriter(Config{
		Level:  "debug",
		Format: "text",
	}, &buf)

	if err != nil {
		t.Fatal(err)
	}

	log.Debug("testing")

	if !strings.Contains(buf.String(), "testing") {
		t.Fatal("text logger failed")
	}
}

func TestInvalidLevel(t *testing.T) {

	_, err := New(Config{
		Level:  "banana",
		Format: "json",
	})

	if err == nil {
		t.Fatal("expected invalid level error")
	}
}
