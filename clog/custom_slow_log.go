package clog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

type CustomConsoleWriter struct {
	Out      io.Writer
	CallSkip int
}

func NewCustomConsoleWriter(callSkip int, options ...func(w *CustomConsoleWriter)) CustomConsoleWriter {
	w := CustomConsoleWriter{
		Out:      os.Stdout,
		CallSkip: callSkip,
	}

	for _, opt := range options {
		opt(&w)
	}

	return w
}

func (w CustomConsoleWriter) Write(p []byte) (n int, err error) {
	var evt map[string]interface{}
	err = json.Unmarshal(p, &evt)
	if err != nil {
		return n, fmt.Errorf("cannot decode event: %s", err)
	}

	var caller string
	if _, file, line, ok := runtime.Caller(w.CallSkip); ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	var prefix map[string]interface{} = nil
	ev := evt[`prefix`]
	if ev != nil {
		prefix = ev.(map[string]interface{})
	}

	// Create a custom log structure with the specific order
	logStructure := CustomLogStructure{
		Level:     fmt.Sprint(evt["level"]),
		Message:   fmt.Sprint(evt["message"]),
		Meta:      prefix,
		Timestamp: time.Now().Format(time.RFC3339),
		Caller:    caller,
	}

	// Serialize the custom log structure to JSON
	output, err := json.Marshal(logStructure)
	if err != nil {
		return n, fmt.Errorf("cannot marshal event: %s", err)
	}

	w.Out.Write(output)
	w.Out.Write([]byte("\n"))
	return len(p), nil
}

// CustomLogStructure defines the log structure with specific field order
type CustomLogStructure struct {
	Level     string                 `json:"level"`
	Message   string                 `json:"msg"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
	Timestamp string                 `json:"timestamp"`
	Caller    string                 `json:"caller"`
}
