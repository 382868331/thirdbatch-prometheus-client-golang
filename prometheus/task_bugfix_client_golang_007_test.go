package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang007SourceContract(t *testing.T) {
    source, err := os.ReadFile("go_collector_latest.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "sampleMap[d.Name] = &sampleBuf[len(sampleBuf)-1]") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "sampleMap[d.Name] = &sampleBuf[len(sampleBuf)+ 1]") {
        t.Fatalf("mutated source contract is still present")
    }
}
