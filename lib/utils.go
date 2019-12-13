package lib

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// GetString ...
func GetString(path string) (string, error) {
	// f := LoadInput(path)

	bits, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// l, err := ReadLine(f)
	// if err != nil {
	// 	return "", err
	// }
	return string(bits), nil
}

// LoadInput returns an *os.File for use with LoopOverLines. If the file
// cannot be found it will print the error and then call os.Exit(1)
func LoadInput(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("unable to open input: %v\n", err)
		os.Exit(1)
	}
	return f
}

// LoopOverLines takes a file, and then calls the provided file on each line
func LoopOverLines(file *os.File, fn func(line []byte) error) error {
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	for ; err == nil; line, _, err = r.ReadLine() {
		x := fn(line)
		if x != nil {
			return err
			// fmt.Printf("\n\ngot error: %v\n", x)
			// os.Exit(1)
		}
	}

	if err == io.EOF {
		return nil
	}

	return err
}

// LoadAndLoop ...
func LoadAndLoop(path string, fn func(ln string) error) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = LoopOverLines(f, func(in []byte) error {
		return fn(string(in))
	})

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// ReadLine reads a single line from the input file
func ReadLine(file *os.File) ([]byte, error) {
	r := bufio.NewReader(file)
	fi, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}
	out := make([]byte, fi.Size())
	_, err = r.Read(out)
	return out, err
}
