package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
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

	phaseInputs := buildAllPhases()
	phaseResult := 0

	for _, phases := range phaseInputs {
		output := execute(intInstr, phases, 0)
		fmt.Printf("Output of %v is %v\n", phases, output)
		if output > phaseResult {
			phaseResult = output
		}
	}
	// phaseResult = execute(intInstr, [5]int{9, 8, 7, 6, 5}, 0)

	fmt.Printf("Highest output is %v\n", phaseResult)
}

func execute(intInstr []int, phases [5]int, input int) int {
	add := 1
	mult := 2
	store := 3
	output := 4
	jumpIfTrue := 5
	jumpIfFalse := 6
	lessThan := 7
	equals := 8
	halt := 99

	phaseUsed := [5]bool{false, false, false, false, false}
	phaseMem := getPhaseInstr(intInstr, 5)

	pointer := [5]int{0, 0, 0, 0, 0}
	instJump := [5]int{0, 0, 0, 0, 0}

	curMem := phaseMem[0]
	curPointer := pointer[0]
	curInstJump := instJump[0]

	phasePointer := 0
	exit := false
	returnVal := 0

	for curPointer < len(curMem) && !exit {
		mode, opcode := processInstr(curMem[curPointer])

		switch opcode {
		case add:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			arg2 := getArg(mode[1], curMem, curMem[curPointer+2])
			sum := arg1 + arg2
			curMem[curMem[curPointer+3]] = sum
			curInstJump = 4
		case mult:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			arg2 := getArg(mode[1], curMem, curMem[curPointer+2])
			prod := arg1 * arg2
			curMem[curMem[curPointer+3]] = prod
			curInstJump = 4
		case store:
			if phaseUsed[phasePointer] != true {
				curMem[curMem[curPointer+1]] = phases[phasePointer]
				phaseUsed[phasePointer] = true
			} else {
				curMem[curMem[curPointer+1]] = input
			}
			curInstJump = 2
		case output:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			returnVal = arg1
			input = arg1
			curInstJump = 2

			// save current memory
			instJump[phasePointer] = curInstJump
			pointer[phasePointer] = curPointer + curInstJump
			phaseMem[phasePointer] = curMem

			// update phase ref
			if phasePointer == 4 {
				phasePointer = 0
			} else {
				phasePointer++
			}

			// change memory spaces
			curMem = phaseMem[phasePointer]
			curInstJump = instJump[phasePointer]
			curPointer = pointer[phasePointer]

			// dont process anything else in the loop
			continue
		case jumpIfTrue:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			arg2 := getArg(mode[1], curMem, curMem[curPointer+2])
			if arg1 != 0 {
				curPointer = arg2
				curInstJump = 0
			} else {
				curInstJump = 3
			}
		case jumpIfFalse:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			arg2 := getArg(mode[1], curMem, curMem[curPointer+2])
			if arg1 == 0 {
				curPointer = arg2
				curInstJump = 0
			} else {
				curInstJump = 3
			}
		case lessThan:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			arg2 := getArg(mode[1], curMem, curMem[curPointer+2])
			if arg1 < arg2 {
				curMem[curMem[curPointer+3]] = 1
			} else {
				curMem[curMem[curPointer+3]] = 0
			}
			curInstJump = 4
		case equals:
			arg1 := getArg(mode[2], curMem, curMem[curPointer+1])
			arg2 := getArg(mode[1], curMem, curMem[curPointer+2])
			if arg1 == arg2 {
				curMem[curMem[curPointer+3]] = 1
			} else {
				curMem[curMem[curPointer+3]] = 0
			}
			curInstJump = 4
		case halt:
			exit = true
		}

		curPointer += curInstJump
	}
	return returnVal

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

func buildAllPhases() [][5]int {
	allPhases := make([][5]int, 0)
	for i := 5; i < 10; i++ {
		for j := 5; j < 10; j++ {
			if i == j {
				continue
			}
			for k := 5; k < 10; k++ {
				if k == j || k == i {
					continue
				}
				for l := 5; l < 10; l++ {
					if l == k || l == j || l == i {
						continue
					}
					for m := 5; m < 10; m++ {
						if m != l && m != k && m != j && m != i {
							allPhases = append(allPhases, [5]int{i, j, k, l, m})
						}
					}
				}
			}
		}
	}
	return allPhases
}

func getPhaseInstr(instr []int, count int) [][]int {
	phaseMem := make([][]int, count)

	for _, val := range instr {
		for i := 0; i < count; i++ {
			phaseMem[i] = append(phaseMem[i], val)
		}
	}

	return phaseMem
}
