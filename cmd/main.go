package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/kmtr/atai"
	"github.com/kmtr/textslicer"
)

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

	scanner := bufio.NewScanner(f)
	slicer := &textslicer.XlsxSlicer{
		ChunkProcessor: &textslicer.ChunkPrinter{
			NameMaker: &textslicer.XlsxNameMaker{
				Prefix: f.Name(),
			},
		},
	}
	if err := slicer.Slice(n, scanner); err != nil {
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
