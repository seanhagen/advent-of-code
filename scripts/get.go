package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/gocolly/colly"
)

var mgoTmpl = `package main

import (
  "fmt"
  "os"
)

/*
{{.}}
*/

func main(){
  fmt.Printf("nope!\n")
  os.Exit(1)
}
`

type tmplContent struct {
	Desc string
}

func get() {
	if len(os.Args) < 4 {
		fmt.Printf("year and day required!\n")
		os.Exit(1)
	}

	c := colly.NewCollector()

	year := strings.TrimSpace(os.Args[2])
	day := strings.TrimSpace(os.Args[3])

	if year == "" {
		year = "2015"
	}

	if day == "" {
		day = "1"
	}

	var err error
	tmpl := template.New("t")
	tmpl, err = tmpl.Parse(mgoTmpl)
	if err != nil {
		fmt.Printf("unable to parse output template: %v\n", err)
		os.Exit(1)
	}

	desc := ""
	wg := &sync.WaitGroup{}

	c.OnHTML("article.day-desc", func(e *colly.HTMLElement) {
		e.ForEach("article.day-desc > h2, article.day-desc > p,article.day-desc >ul, article.day-desc > pre", func(i int, e *colly.HTMLElement) {

			switch e.Name {
			case "ul":
				e.ForEach("li", func(_ int, e *colly.HTMLElement) {
					desc = fmt.Sprintf("%v\n\t- %v\n", desc, e.Text)
				})

			case "pre":
				desc = fmt.Sprintf("%v\n%v\n", desc, e.Text)

			default:
				desc = fmt.Sprintf("%v\n%v\n", desc, colw(e.Text, "", 100))
			}
		})
		wg.Done()
	})

	wg.Add(1)
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v", year, day)
	c.Visit(url)

	dt := fmt.Sprintf("%v", day)
	if len(dt) == 1 {
		dt = "0" + dt
	}

	wg.Wait()

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, desc)
	if err != nil {
		fmt.Printf("unable to render template: %v\n", err)
		os.Exit(1)
	}

	path := fmt.Sprintf("../%v/day%v/part1/main.go", year, dt)
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

const space = ' '

func colw(input, ls string, col int) string {
	if len(input) < col {
		return input
	}

	bits := []string{}
	idx := col
	s, e := 0, 0

	for ; idx < len(input); idx += col {
		if input[idx] == space {
			// it's a space, go forward to first non-space character
			// ie, if string is ".... ...   ..."
			// and idx is here -----------^
			// go forward to here ----------^
			for {
				idx++
				if input[idx] != space {
					break
				}
			}

		} else {
			// it's a character, go back to first character in string from idx
			// ie, if string is "... .. ..."
			// and idx is here ----------^
			// go back to here ---------^
			for idx > 0 {
				idx--
				if input[idx] == space {
					idx++
					break
				}
			}
		}

		e = idx - 1
		bits = append(bits, ls+input[s:e])
		s = idx
	}

	bits = append(bits, input[s:])
	return strings.Join(bits, "\n")
}
