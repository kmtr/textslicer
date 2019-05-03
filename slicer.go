package textslicer

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type ChunkProcessor interface {
	Proc(chunk *list.List)
}

type ChunkPrinter struct {
	NameMaker NameMaker
}

func (cp *ChunkPrinter) Proc(chunk *list.List) {
	fmt.Println(cp.NameMaker.Name())
	for e := chunk.Front(); e != nil; e = e.Next() {
		fmt.Fprintln(os.Stdout, e.Value)
	}
}

type NameMaker interface {
	Name() string
}

type XlsxNameMaker struct {
	Prefix  string
	counter int
}

func (xnm *XlsxNameMaker) Name() string {
	xnm.counter++
	return fmt.Sprintf("%s-%03d.%s", xnm.Prefix, xnm.counter, "xlsx")
}

type Slicer interface {
	Slice(n int, scanner *bufio.Scanner) error
}

type XlsxSlicer struct {
	ChunkProcessor ChunkProcessor
}

func (xlss *XlsxSlicer) Slice(n int, scanner *bufio.Scanner) error {
	return slice(n, scanner, xlss.ChunkProcessor)
}

func slice(n int, scanner *bufio.Scanner, cp ChunkProcessor) error {
	var lineList *list.List
	i := 0
	for scanner.Scan() {
		i++
		if lineList == nil {
			lineList = list.New()
		}
		line := scanner.Text()
		if line != "" {
			lineList.PushBack(line)
		}
		if n == lineList.Len() {
			cp.Proc(lineList)
			lineList = nil
		}
	}
	if lineList.Len() > 0 {
		cp.Proc(lineList)
		lineList = nil
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
