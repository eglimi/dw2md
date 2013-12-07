package main

import (
	"flag"
	"fmt"
	"github.com/eglimi/dw2md/scanner"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	infile := flag.String("i", "", "The input file. stdin if not set.")
	outfile := flag.String("o", "", "The output file. stdout if not set.")
	help := flag.Bool("help", false, "Show this help.")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var err error

	// In / Output
	var inReader io.ReadCloser
	var outWriter io.WriteCloser
	if *infile != "" {
		// use a file for input
		absIn, err := filepath.Abs(*infile)
		if err != nil {
			fmt.Printf("Could not find infile %s. Error %v\n", *infile, err)
			os.Exit(1)
		}
		inReader, err = os.Open(absIn)
		if err != nil {
			fmt.Printf("Could not open file %v. Error %v\n", absIn, err)
		}
		defer inReader.Close()
	} else {
		// use stdin
		inReader = os.Stdin
	}
	if *outfile != "" {
		// use a file for output
		absOut, err := filepath.Abs(*outfile)
		if err != nil {
			fmt.Printf("Could not find outfile %s. Error %v\n", *outfile, err)
			os.Exit(1)
		}
		fileMod := os.O_RDWR | os.O_CREATE | os.O_TRUNC;
		outWriter, err = os.OpenFile(absOut, fileMod, 0666)
		if err != nil {
			fmt.Printf("Could not open file %v. Error %v\n", absOut, err)
		}
		defer outWriter.Close()
	} else {
		// use stdout
		outWriter = os.Stdout
	}

	in, err := ioutil.ReadAll(inReader)
	if err != nil {
		fmt.Printf("Could not read from input. Error: %v\n", err)
		os.Exit(1)
	}

	out := scanner.ConvertDoc(in)

	_, err = outWriter.Write(out)
	if err != nil {
		fmt.Printf("Could not write to output. Error: %v\n", err)
		os.Exit(1)
	}
}
