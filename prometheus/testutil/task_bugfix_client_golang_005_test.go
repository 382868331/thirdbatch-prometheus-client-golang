package testutil

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang005SourceContract(t *testing.T) {
    source, err := os.ReadFile("testutil.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
