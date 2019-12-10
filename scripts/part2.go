package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func part2() {
	if len(os.Args) < 4 {
		fmt.Printf("year and day required!\n")
		os.Exit(1)
	}

	year := strings.TrimSpace(os.Args[2])
	day := strings.TrimSpace(os.Args[3])
	if year == "" {
		year = "2015"
	}

	if day == "" {
		day = "1"
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | aoc")
		os.Exit(1)
	}

	data := ""

	scanner := bufio.NewReader(os.Stdin)

	for {
		str, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error while reading: %v\n", err)
			os.Exit(1)
		}
		data = data + colw(str, "", 100)
	}

	tmpl := template.New("t")
	tmpl, err = tmpl.Parse(mgoTmpl)
	if err != nil {
		fmt.Printf("unable to parse output template: %v\n", err)
		os.Exit(1)
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, data)
	if err != nil {
		fmt.Printf("unable to render template: %v\n", err)
		os.Exit(1)
	}

	dt := fmt.Sprintf("%v", day)
	if len(dt) == 1 {
		dt = "0" + dt
	}

	path := fmt.Sprintf("../%v/day%v/part2/main.go", year, dt)
	fmt.Printf("output to file: %v\n", path)

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("unable to open file: %v\n")
		os.Exit(1)
	}

	n, err := f.WriteString(buf.String())
	if err != nil {
		fmt.Printf("unable to write to file: %v\n")
		os.Exit(1)
	}

	fmt.Printf("wrote %v bytes to %v!\n", n, path)
}
