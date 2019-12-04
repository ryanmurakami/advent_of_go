package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	instJump := 4
	add := 1
	mult := 2
	halt := 99

	bs, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(bs)
	instr := strings.Split(input, ",")

	// convert instructions to ints
	intInstr := make([]int, len(instr))

	for i, strVal := range instr {
		intInstr[i], _ = strconv.Atoi(strVal)
	}

	// replace instructions
	intInstr[1] = 57
	intInstr[2] = 41 // adding to end

	pointer := 0

	for pointer < len(intInstr) {
		if intInstr[pointer] == add {
			sum := intInstr[intInstr[pointer+1]] + intInstr[intInstr[pointer+2]]
			intInstr[intInstr[pointer+3]] = sum
		} else if intInstr[pointer] == mult {
			prod := intInstr[intInstr[pointer+1]] * intInstr[intInstr[pointer+2]]
			intInstr[intInstr[pointer+3]] = prod
		} else if intInstr[pointer] == halt {
			break
		}
		pointer += instJump
	}

	fmt.Println(intInstr[0])
}
