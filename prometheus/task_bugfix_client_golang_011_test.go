package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang011SourceContract(t *testing.T) {
    source, err := os.ReadFile("wrap.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := r.Register(c); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
