package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/karrick/gotb"
)

func main() {
	optNumber := flag.Int("n", 0, "display N final lines")
	flag.Parse()

	if err := notTail(*optNumber, os.Stdin, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

// notTail copies all but the final num lines from io.Reader to io.Writer.
func notTail(num int, r io.Reader, w io.Writer) error {
	cb, err := gotb.NewStrings(num)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line, ok := cb.QueueDequeue(scanner.Text())
		if !ok {
			continue
		}

		if _, err = fmt.Fprintln(w, line); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
