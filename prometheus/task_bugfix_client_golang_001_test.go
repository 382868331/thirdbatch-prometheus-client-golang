package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang001SourceContract(t *testing.T) {
    source, err := os.ReadFile("registry.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return true") {
        t.Fatalf("expected source contract is missing")
    }
}
