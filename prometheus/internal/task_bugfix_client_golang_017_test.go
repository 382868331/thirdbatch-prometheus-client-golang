package internal

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixClientGolang017SourceContract(t *testing.T) {
    source, err := os.ReadFile("difflib.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "codes[len(codes)-1] = OpCode{c.Tag, i1, minInt(i2, i1+n), j1, minInt(j2, j1+n)}") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "codes[len(codes)+ 1] = OpCode{c.Tag, i1, minInt(i2, i1+n), j1, minInt(j2, j1+n)}") {
        t.Fatalf("mutated source contract is still present")
    }
}
