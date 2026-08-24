package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang003SourceContract(t *testing.T) {
    source, err := os.ReadFile("registry.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(r.collectorsByID) == 0 && len(r.uncheckedCollectors) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
