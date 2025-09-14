//go:build linux

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"time"
)

const (
	version     = "0.0.1\n"
	usageString = `Usage: passgen [flags]

Flags:
	-l, --length         specify the length of password
	-e  --easy-read      avoid ambiguous characters like l, 1, O, and 0
	-h, --help           print help information
    -v, --version        print version

Examples:
  passgen --length 64
  passgen -l 64
  passgen -l 64 --ease-read

`
)

// flags
var (
	fLength   int
	fEasyRead bool
	fVersion  bool
	fHelp     bool
)

const (
	randomStream = "/dev/random"
)

var (
	avoidAmbiguousChars = []byte{48, 49, 79, 108}
)

func main() {
	flag.IntVar(&fLength, "length", 0, "specify the length of password")
	flag.IntVar(&fLength, "l", 0, "specify the length of password")
	flag.BoolVar(&fEasyRead, "easy-read", false, "avoid ambiguous characters like l, 1, O, and 0")
	flag.BoolVar(&fEasyRead, "e", false, "avoid ambiguous characters like l, 1, O, and 0")
	flag.BoolVar(&fVersion, "version", false, "print version")
	flag.BoolVar(&fVersion, "v", false, "print version")
	flag.BoolVar(&fHelp, "help", false, "print help")
	flag.BoolVar(&fHelp, "h", false, "print help")

	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usageString)
		os.Exit(0)
	}
	flag.Parse()

	run()
}

func run() {
	switch {
	case fHelp:
		printUsage()
		os.Exit(0)
	case fVersion:
		printVersion()
		os.Exit(0)
	case fLength <= 0:
		fmt.Println("Length should be greater than 0")
		os.Exit(-1)
	case fLength > 0:
		passwd, err := generatePassword(fLength)

		if err != nil {
			fmt.Fprintf(os.Stdout, "Couldn't generate password, error: %s", err.Error())
			os.Exit(-1)
		}

		fmt.Println(passwd)
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "flag provided but not defined %s \n", flag.Args())
		printUsage()
		os.Exit(-1)
	}
}

func printVersion() {
	fmt.Fprintf(os.Stdout, version)
}

func printUsage() {
	fmt.Fprintf(os.Stdout, usageString)
}

func generatePassword(length int) (string, error) {

	start := time.Now()
	defer func() {
		endTime := time.Since(start)

		fmt.Println(endTime)
	}()

	randomFile, err := os.OpenFile(randomStream, os.O_RDONLY, os.ModeDevice)

	if err != nil {
		return "", err
	}

	bufReader := bufio.NewReader(randomFile)

	buf := make([]byte, 0)

	var i int

	for {

		if i == length {
			break
		}

		currentByte, err := bufReader.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}

		if err != nil {
			break
		}

		currentByteInRange := (currentByte % (126 - 33)) + 33

		if fEasyRead && slices.Contains(avoidAmbiguousChars, currentByteInRange) {
			continue
		}

		buf = append(buf, currentByteInRange)

		i++
	}

	return string(buf), nil
}
