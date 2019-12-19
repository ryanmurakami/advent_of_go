package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	instr := readInput("input")

	process(instr)
}

func readInput(filename string) []int {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return make([]int, 0)
	}

	input := string(bs)
	instr := strings.Split(input, ",")

	// convert instructions to ints
	intInstr := make([]int, len(instr))

	for i, strVal := range instr {
		intInstr[i], _ = strconv.Atoi(strVal)
	}

	return intInstr
}
