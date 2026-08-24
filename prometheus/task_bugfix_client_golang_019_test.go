package prometheus

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang019SourceContract(t *testing.T) {
    source, err := os.ReadFile("wrap.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "func (c *wrappingCollector) Describe(ch chan<- *Desc) {") {
        t.Fatalf("expected source contract is missing")
    }
}
