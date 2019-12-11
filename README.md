Advent of Code
==============

My solutions for each year's Advent of Code. Each year has it's own directory,
and each day has it's own folder under that. If I've solved a problem using more
than one language, then that day will have a folder for each language. Mostly
it's all in [Go](https://golang.org), though.


## Helper Script

In `scripts`, compile helper with `go build . -out aoc`.

### Get Part 1

In `scripts`, run `./aoc get <year> <day>` to get the part 1 description. Will
create a `main.go` file in `<year>/day<day>/part1` ( day will be 0 padded if
less than 10 ).

### Part 2

In `scripts`, run `xclip -out | ./aoc part2 <year> <day>` to get the part 2
description. Creates the same type of main.go that `get` creates, it just does
it in the part2 folder for the given day.
