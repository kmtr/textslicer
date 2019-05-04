package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/kmtr/atai"
	"github.com/kmtr/textslicer"
	"github.com/tealeg/xlsx"
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
	xlsxFile := xlsx.NewFile()
	cp := textslicer.NewXlsxPrinter(xlsxFile)
	/*
		&textslicer.ChunkPrinter{
			NameMaker: &textslicer.XlsxNameMaker{
				Prefix: f.Name(),
			},
		},
	*/
	slicer := &textslicer.XlsxSlicer{
		ChunkProcessor: cp,
	}
	if err := slicer.Slice(n, scanner); err != nil {
		return err
	}
	xlsxFile.Save("./sample.xlsx")
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
