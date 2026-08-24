package api

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang015SourceContract(t *testing.T) {
    source, err := os.ReadFile("client.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer close(done)") {
        t.Fatalf("expected source contract is missing")
    }
}
