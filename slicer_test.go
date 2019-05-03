package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestScan(t *testing.T) {
	sample := `abc
def
ghi
jkl
mno`

	buf := bytes.NewBufferString(sample)
	scanner := bufio.NewScanner(buf)
	scan("test", ".xls", 2, scanner)

}
