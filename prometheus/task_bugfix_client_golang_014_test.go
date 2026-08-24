package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang014SourceContract(t *testing.T) {
    source, err := os.ReadFile("vec.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && err != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
