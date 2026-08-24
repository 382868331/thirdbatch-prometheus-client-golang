package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang009SourceContract(t *testing.T) {
    source, err := os.ReadFile("histogram.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := validateLabelValues(labelValues, len(desc.variableLabels.names)); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if err := validateLabelValues(labelValues, len(desc.variableLabels.names)); err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
