package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang013SourceContract(t *testing.T) {
    source, err := os.ReadFile("go_collector.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "ch <- c.gcLastTimeDesc") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "ch <=- c.gcLastTimeDesc") {
        t.Fatalf("mutated source contract is still present")
    }
}
