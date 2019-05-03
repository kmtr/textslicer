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
		ChunkProcessor: &ChunkPrinter{
			NameMaker: &XlsxNameMaker{
				Prefix: "test",
			},
		},
	}
	buf := bytes.NewBufferString(sample)
	scanner := bufio.NewScanner(buf)
	slicer.Slice(2, scanner)
}
