package graphite

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang002SourceContract(t *testing.T) {
    source, err := os.ReadFile("bridge.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if c.Gatherer == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && c.Gatherer == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
