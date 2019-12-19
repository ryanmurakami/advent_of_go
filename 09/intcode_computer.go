package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func process(intInstr []int) {
	instJump := 4
	add := 1
	mult := 2
	store := 3
	output := 4
	jumpIfTrue := 5
	jumpIfFalse := 6
	lessThan := 7
	equals := 8
	adjust := 9
	halt := 99

	pointer := 0
	relBase := 0
	exit := false

	for pointer < len(intInstr) && !exit {
		mode, opcode := processInstr(intInstr[pointer])

		switch opcode {
		case add:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2], relBase)
			sum := arg1 + arg2
			intInstr = set(mode[0], intInstr, intInstr[pointer+3], sum, relBase)
			instJump = 4
		case mult:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2], relBase)
			prod := arg1 * arg2
			intInstr = set(mode[0], intInstr, intInstr[pointer+3], prod, relBase)
			instJump = 4
		case store:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter code:")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			intText, _ := strconv.Atoi(text)
			intInstr = set(mode[2], intInstr, intInstr[pointer+1], intText, relBase)
			instJump = 2
		case output:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			fmt.Printf("Output: %d\n", arg1)
			instJump = 2
		case jumpIfTrue:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2], relBase)
			if arg1 != 0 {
				pointer = arg2
				instJump = 0
			} else {
				instJump = 3
			}
		case jumpIfFalse:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2], relBase)
			if arg1 == 0 {
				pointer = arg2
				instJump = 0
			} else {
				instJump = 3
			}
		case lessThan:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2], relBase)
			if arg1 < arg2 {
				intInstr = set(mode[0], intInstr, intInstr[pointer+3], 1, relBase)
			} else {
				intInstr = set(mode[0], intInstr, intInstr[pointer+3], 0, relBase)
			}
			instJump = 4
		case equals:
			arg1 := getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			arg2 := getArg(mode[1], intInstr, intInstr[pointer+2], relBase)
			if arg1 == arg2 {
				intInstr = set(mode[0], intInstr, intInstr[pointer+3], 1, relBase)
			} else {
				intInstr = set(mode[0], intInstr, intInstr[pointer+3], 0, relBase)
			}
			instJump = 4
		case adjust:
			relBase += getArg(mode[2], intInstr, intInstr[pointer+1], relBase)
			instJump = 2
		case halt:
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

func getArg(mode int, instr []int, pos int, relBase int) int {
	if mode == 0 {
		if pos >= len(instr) {
			return 0
		}
		return instr[pos]
	} else if mode == 1 {
		return pos
	}
	if relBase+pos >= len(instr) {
		return 0
	}
	return instr[relBase+pos]
}

func set(mode int, instr []int, pos int, val int, relBase int) []int {
	instrLen := len(instr)
	var finalPos int
	if mode == 2 {
		finalPos = relBase + pos
	} else {
		finalPos = pos
	}

	if finalPos >= instrLen {
		newInstr := make([]int, finalPos+1)
		copy(newInstr, instr)
		newInstr[finalPos] = val
		return newInstr
	}

	instr[finalPos] = val
	return instr
}
