package main

import (
	"flag"
	"fmt"
)

var (
	n bool
	h bool

	SPACE   = " "
	NEWLINE = "\n"
)

func main() {
	flag.BoolVar(&n, "n", false, "do not output the trailing newline")
	flag.BoolVar(&h, "h", false, "print help")

	flag.Parse()

	if h {
		flag.PrintDefaults()
		return
	}

	argLen := flag.NArg() - 1
	outLine := ""
	for idx, arg := range flag.Args() {
		outLine += arg
		if idx < argLen {
			outLine += SPACE
		}
	}
	fmt.Printf("%s", outLine)

	// nオプションが有効であれば改行は出力しない
	if !n {
		fmt.Printf(NEWLINE)
	}
}
