package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisClientGolang004SourceContract(t *testing.T) {
    source, err := os.ReadFile("go_collector.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "ch <- MustNewConstMetric(c.goInfoDesc, GaugeValue, 1)") {
        t.Fatalf("expected source contract is missing")
    }
}
