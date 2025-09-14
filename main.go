package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version     = "0.0.1\n"
	usageString = `Usage: passgen [flags]

Flags:
	-f, --length         specify the length of password
	-h, --help           print help information
    -v, --version        print version

Examples:
  passgen comma.csv              | c2j | jq
  passgen semicolon.csv          | c2j --delimiter ";" | jq
  passgen csv_without_header.csv | c2j --no-header | jq

`
)

// flags
var (
	fLength  int
	fVersion bool
	fHelp    bool
)

func main() {
	flag.IntVar(&fLength, "length", 0, "specify the length of password")
	flag.IntVar(&fLength, "l", 0, "specify the length of password")
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
		fmt.Println(os.Stdout, "Length should be greater than 0 %s")
		printUsage()
		os.Exit(-1)
	case fLength > 0:
		passwd, err := generatePassword(fLength)

		if err != nil {
			fmt.Fprintf(os.Stdout, "flag provided but not defined %s \n", flag.Args())
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

	randomFile, err := os.OpenFile("/dev/urandom", os.O_RDONLY, os.ModeDevice)

	if err != nil {
		return "", err
	}

	buf := make([]byte, length)

	_, err = randomFile.ReadAt(buf, 0)

	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func debug() {

	debug := os.Getenv("DEBUG")

	if len(debug) > 0 {
		fmt.Fprintf(os.Stdout, "%s", "")
	}
}
