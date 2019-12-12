package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	instJump := 4
	ADD := 1
	MULT := 2
	STORE := 3
	OUTPUT := 4
	JUMP_IF_TRUE := 5
	JUMP_IF_FALSE := 6
	LESS_THAN := 7
	EQUALS := 8
	HALT := 99

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

	pointer := 0
	exit := false

	for pointer < len(intInstr) && !exit {
		mode, opcode := processInstr(intInstr[pointer])

		switch opcode {
		case ADD:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2])
			sum := arg1 + arg2
			intInstr[intInstr[pointer+3]] = sum
			instJump = 4
		case MULT:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2])
			prod := arg1 * arg2
			intInstr[intInstr[pointer+3]] = prod
			instJump = 4
		case STORE:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter code:")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			intText, _ := strconv.Atoi(text)
			intInstr[intInstr[pointer+1]] = intText
			instJump = 2
		case OUTPUT:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			fmt.Printf("Output: %d\n", arg1)
			instJump = 2
		case JUMP_IF_TRUE:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2])
			if arg1 != 0 {
				pointer = arg2
				instJump = 0
			} else {
				instJump = 3
			}
		case JUMP_IF_FALSE:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2])
			if arg1 == 0 {
				pointer = arg2
				instJump = 0
			} else {
				instJump = 3
			}
		case LESS_THAN:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2])
			if arg1 < arg2 {
				intInstr[intInstr[pointer+3]] = 1
			} else {
				intInstr[intInstr[pointer+3]] = 0
			}
			instJump = 4
		case EQUALS:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1])
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2])
			if arg1 == arg2 {
				intInstr[intInstr[pointer+3]] = 1
			} else {
				intInstr[intInstr[pointer+3]] = 0
			}
			instJump = 4
		case HALT:
			exit = true
		}

		pointer += instJump
	}
}

func processInstr(instr int) ([]int, int) {
	mode := make([]int, 3)

	opcode := instr % 100
	mode[2] = instr % 1000 / 100
	mode[1] = instr % 10000 / 1000
	mode[0] = instr % 100000 / 10000

	return mode, opcode
}

func getArg(mode int, instr []int, pos int) int {
	if mode == 0 {
		return instr[pos]
	}
	return pos
}
