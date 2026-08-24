package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang006SourceContract(t *testing.T) {
    source, err := os.ReadFile("process_collector_darwin.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if addressSpace, err := getSoftLimit(syscall.RLIMIT_AS); err == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
