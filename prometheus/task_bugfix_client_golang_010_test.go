package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang010SourceContract(t *testing.T) {
    source, err := os.ReadFile("counter.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
