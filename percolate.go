package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/echinmay/unionfind"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	grid_x   int
	grid_dim int
	upper    int
	lower    int
)

func isvalidinput(grid_x, p, q int) bool {
	if p <= 0 || q <= 0 || p > grid_x || q > grid_x {
		return false
	}
	return true
}

func convertxytoarridx(grid_x, row, col int) int {
	idx := ((row-1)*grid_x + (col - 1))
	return idx
}

func open(uf *unionfind.UnionFind, opened []bool, x, y int) {

	entry := convertxytoarridx(grid_x, x, y)
	opened[entry] = true

	if x == 1 {
		// This is the top row.
		uf.Union(upper-1, entry)
	} else {
		upentry := convertxytoarridx(grid_x, x-1, y)
		if opened[upentry] {
			uf.Union(upentry, entry)
		}
	}

	if x == grid_x {
		// This is the bottom row
		uf.Union(lower-1, entry)
	} else {
		downentry := convertxytoarridx(grid_x, x+1, y)
		if opened[downentry] {
			uf.Union(downentry, entry)
		}
	}

	if !(y == 1) {
		leftentry := convertxytoarridx(grid_x, x, y-1)
		if opened[leftentry] {
			uf.Union(leftentry, entry)
		}
	}

	if !(y == grid_x) {
		rightentry := convertxytoarridx(grid_x, x, y+1)
		if opened[rightentry] {
			uf.Union(rightentry, entry)
		}
	}

	return
}

func main() {

	var inputfile string

	// Get the input file if sent as command line argument or use default
	flag.StringVar(&inputfile, "input", "input.txt", "<NAME OF INPUT FILE>")
	flag.Parse()

	//fmt.Println(inputfile)
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	//fmt.Println(line)
	inputs := strings.Fields(line)
	//fmt.Println(inputs)

	grid_x, _ = strconv.Atoi(inputs[0])
	//	fmt.Println("Grid dimension", grid_x)
	grid_dim = grid_x * grid_x
	upper = grid_dim + 1
	lower = grid_dim + 2

	uf := unionfind.New(lower)
	opened := make([]bool, lower)
	for entry := range opened {
		opened[entry] = false
	}

	//	inputs = inputs[1:]
	//for i := 0; i < len(inputs); i = i + 2 {
	linenum := 0
	for scanner.Scan() {
		line = scanner.Text()
		//fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		inputs = strings.Fields(line)
		linenum ++;
		//fmt.Println("---")
		x, err := strconv.Atoi(inputs[0])
		if err != nil {
			log.Fatal(linenum, err)
		}
		y, err := strconv.Atoi(inputs[1])
		if err != nil {
			log.Fatal(linenum, err)
		}
		if !isvalidinput(grid_x, x, y) {
			continue
		}

		open(uf, opened, x, y)
	}

	//fmt.Println(opened)
	fmt.Println(uf.Connected(upper-1, lower-1))
	draw(uf, grid_x, opened)
	//fmt.Println(uf)
}

func draw(uf *unionfind.UnionFind, grid_x int, opened []bool) {
	//y := 0

	for y := 0; y < (grid_x*grid_x); y++ {
		//for y:=0; y < 5; y++ {

		if y%grid_x == 0 {
			fmt.Print("\n")
		}
		if opened[y] == true {
			fmt.Print("* ")
		} else {
			fmt.Print("- ")
		}
	}
	fmt.Println("")
}
