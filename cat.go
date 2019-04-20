package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	number        bool
	numberNoBlank bool
	showAll       bool
	showEnds      bool
	showTabs      bool
)

const (
	NEWLINE_REPLACE = "$\n"
	TAB_REPLACE     = "^I"
)

func InitCatOptions() {

	flag.BoolVar(&number, "n", false, "number all output lines")
	flag.BoolVar(&numberNoBlank, "b", false, "number noempty output lines, overrides -n")
	flag.BoolVar(&showAll, "A", false, "equivalent to -vET")
	flag.BoolVar(&showEnds, "E", false, "display $ at end of each line")
	flag.BoolVar(&showTabs, "T", false, "display TAB character as ^I")

	if numberNoBlank {
		number = false
	}
	if showAll {
		showEnds = true
		showTabs = true
	}
	flag.Parse()

}

func CreateReader() (*os.File, error) {
	fmt.Println(flag.Arg(1))
	if len(flag.Args()) < 1 || flag.Arg(0) == "-" {
		return os.Stdin, nil
	} else {
		fp, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		return fp, err
	}
}

func main() {

	// オプションの初期化
	InitCatOptions()

	fp, err := CreateReader()
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	line := 1
	for scanner.Scan() {
		nowLine := scanner.Text() + "\n"
		if number {
			fmt.Printf("%10v ", line)
			line += 1
		}
		if numberNoBlank && nowLine != "\n" {
			fmt.Printf("%10v ", line)
			line += 1
		}
		if showEnds {
			nowLine = strings.ReplaceAll(nowLine, "\n", NEWLINE_REPLACE)
		}
		if showTabs {
			nowLine = strings.ReplaceAll(nowLine, "\t", TAB_REPLACE)
		}
		fmt.Printf(nowLine)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
