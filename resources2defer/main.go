package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("/etc/*")
	if err != nil {
		panic(err)
	}

	failCount := 0
	successCount := 0

	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Printf("failed to open %v\n", fname)
			failCount++
			continue
		}

		defer f.Close()

		l, err := readFirstLine(f)
		if err != nil {
			fmt.Printf("failed to read from %v: %v\n", fname, err)
			failCount++
			continue
		}

		fmt.Printf("%v: %v\n", fname, l)
		successCount++
	}

	fmt.Printf("failed files: %v\n", failCount)
	fmt.Printf("successful files: %v\n", successCount)
}

func readFirstLine(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	if s.Scan() {
		return s.Text(), nil
	}

	if s.Err() == io.EOF {
		return "", errors.New("no first line found")
	}

	return "", s.Err()
}
