package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisClientGolang008SourceContract(t *testing.T) {
    source, err := os.ReadFile("process_collector_mem_cgo_darwin.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "ch <- c.rss") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "ch <=- c.rss") {
        t.Fatalf("mutated source contract is still present")
    }
}
