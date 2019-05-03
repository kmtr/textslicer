package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kmtr/atai"
)

type ChunkProcessor interface {
	Proc(chunk *list.List)
}

type ChunkPrinter struct {
	outputName string
}

func (cp *ChunkPrinter) Proc(chunk *list.List) {
	fmt.Println(cp.outputName)
	for e := chunk.Front(); e != nil; e = e.Next() {
		fmt.Fprintln(os.Stdout, e.Value)
	}
}

func makeOutputName(prefix, suffix string, n int) string {
	return fmt.Sprintf("%s-%03d.%s", prefix, n, suffix)
}

func scan(prefix, suffix string, n int, scanner *bufio.Scanner) error {
	var lineList *list.List

	i := 0
	chunkNum := 0
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
			chunkNum++
			cp := &ChunkPrinter{
				outputName: makeOutputName(prefix, suffix, chunkNum),
			}
			cp.Proc(lineList)
			lineList = nil
		}
	}
	if lineList.Len() > 0 {
		chunkNum++
		cp := &ChunkPrinter{
			outputName: makeOutputName(prefix, suffix, chunkNum),
		}
		cp.Proc(lineList)
		lineList = nil
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil

}

func run(source atai.ValueProvider, sliceNum atai.ValueProvider) error {
	n, err := strconv.Atoi(sliceNum())
	if err != nil {
		return err
	}

	f, err := os.Open(source())
	if err != nil {
		return err
	}
	defer f.Close()

	fname := f.Name()
	scanner := bufio.NewScanner(f)
	if err := scan(fname, ".xlsx", n, scanner); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.String("f", "", "source")
	flag.String("n", "100", "slice num")
	flag.Parse()

	source := atai.ValueFromFlag("f")
	sliceNum := atai.ValueFromFlag("n")
	if err := run(source, sliceNum); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
