package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	byteSize      bool
	lineSize      bool
	wordSize      bool
	multibyteSize bool
)

func init() {
	flag.BoolVar(&byteSize, "c", false, "usage tbd")
	flag.BoolVar(&lineSize, "l", false, "usage tbd")
	flag.BoolVar(&wordSize, "w", false, "usage tbd")
	flag.BoolVar(&multibyteSize, "m", false, "usage tbd")
	flag.Parse()
}

func main() {
	var output string
	switch len(flag.Args()) {
	case 0:
		// no args, read from stdin
		// eg: cat test.txt | go run main.go -l
		b := readFromSTDIN()
		output = message(b, "")
	case 1:
		// just on arg, read from that filename
		// eg: go run main.go -l -w test.txt
		filename := flag.Arg(0)
		b, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		output = message(b, filename)
	default:
		// reading error, too much args
		// eg: go run main.go a.txt b.txt
		fmt.Fprintln(os.Stderr, "too much args")
		os.Exit(1)
	}
	fmt.Println(output)
}

func readFromSTDIN() []byte {
	reader := bufio.NewReader(os.Stdin)
	builder := strings.Builder{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		builder.WriteString(line + "\n")
	}
	return []byte(builder.String())
}

func message(b []byte, filename string) string {
	builder := strings.Builder{}
	// you can specify one or more flags, like the real wc
	if lineSize {
		count := strconv.Itoa(lineCount(b))
		builder.WriteString(count + "\t")
	}
	if wordSize {
		count := strconv.Itoa(wordCount(b))
		builder.WriteString(count + "\t")
	}
	if byteSize {
		count := strconv.Itoa(byteCount(b))
		builder.WriteString(count + "\t")
	}
	if multibyteSize {
		count := strconv.Itoa(multibyteCount(b))
		builder.WriteString(count + "\t")
	}
	// if no flag is set, use l,w,c flags as default
	if builder.Len() == 0 {
		l := lineCount(b)
		w := wordCount(b)
		b := byteCount(b)
		builder.WriteString(
			fmt.Sprintf("%d\t%d\t%d", l, w, b),
		)
	}
	return builder.String() + " " + filename
}

func byteCount(in []byte) int {
	return len(in)
}

func lineCount(in []byte) int {
	lines := strings.FieldsFunc(string(in), func(r rune) bool {
		return r == '\n'
	})
	return len(lines)
}

func wordCount(in []byte) int {
	words := strings.Fields(string(in))
	return len(words)
}

func multibyteCount(in []byte) int {
	var counter int
	for range string(in) {
		counter++
	}
	return counter
}
