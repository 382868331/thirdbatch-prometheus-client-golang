package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisClientGolang016SourceContract(t *testing.T) {
    source, err := os.ReadFile("vec.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false") {
        t.Fatalf("expected source contract is missing")
    }
}
