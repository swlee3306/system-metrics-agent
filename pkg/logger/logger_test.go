package logger

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func capture(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	output, flags, prefix := log.Writer(), log.Flags(), log.Prefix()
	log.SetOutput(&buf)
	log.SetFlags(0)
	log.SetPrefix("")
	t.Cleanup(func() { log.SetOutput(output); log.SetFlags(flags); log.SetPrefix(prefix) })
	return &buf
}

func TestLevelsAndCaller(t *testing.T) {
	buf := capture(t)
	for _, tc := range []struct {
		level string
		call  func(string, string, ...any)
	}{{"DEBUG", Debug}, {"INFO", Info}, {"WARN", Warn}, {"ERROR", Error}, {"FATAL", Fatal}} {
		buf.Reset()
		tc.call("unit", "value=%d", 7)
		got := buf.String()
		for _, want := range []string{"LogLevel:(" + tc.level + ")", "Target:(unit)", "LogMessage:(value=7)", "logger_test.go:"} {
			if !strings.Contains(got, want) {
				t.Errorf("%s missing %q", tc.level, want)
			}
		}
		if strings.Contains(got, "/Users/") || strings.Contains(got, "/home/") {
			t.Error("caller includes build-machine path")
		}
	}
}

func TestConcurrentOutput(t *testing.T) {
	buf := capture(t)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) { defer wg.Done(); Info("worker", "item=%d", n) }(i)
	}
	wg.Wait()
	if got := strings.Count(buf.String(), "LogLevel:(INFO)"); got != 100 {
		t.Fatalf("got %d entries, want 100", got)
	}
}

func TestNoPersistenceDependency(t *testing.T) {
	source, err := os.ReadFile("logger.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"pkg/db/model", "gorm.io/", "LogEntries", "RemoveAll", "time.NewTicker"} {
		if strings.Contains(string(source), forbidden) {
			t.Errorf("legacy dependency: %s", forbidden)
		}
	}
	models, err := filepath.Glob("../db/model/*.go")
	if err != nil {
		t.Fatal(err)
	}
	if len(models) != 0 {
		t.Errorf("%d legacy model source files remain", len(models))
	}
}
