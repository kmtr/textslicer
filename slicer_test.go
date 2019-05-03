package textslicer

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

	slicer := &XlsxSlicer{
		NameMaker: new(XlsxNameMaker),
	}
	buf := bytes.NewBufferString(sample)
	scanner := bufio.NewScanner(buf)
	slicer.Slice(2, scanner)
}
