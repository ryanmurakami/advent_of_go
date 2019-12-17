package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	instr := readInput("input")
	image := buildLayers(instr)

	var finalImage [6][25]int

	// initialize finalImage

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			finalImage[i][j] = 2
		}
	}

	for _, layer := range image {
		for i := 0; i < 6; i++ {
			for j := 0; j < 25; j++ {
				if finalImage[i][j] == 2 {
					finalImage[i][j] = layer[i][j]
				}
			}
		}
	}

	// print finalImage
	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			if finalImage[i][j] == 0 {
				fmt.Printf(" ")
			} else if finalImage[i][j] == 1 {
				fmt.Printf("#")
			} else if finalImage[i][j] == 2 {
				fmt.Printf("X")
			}
		}
		fmt.Printf("\n")
	}
}

func readInput(filename string) []int {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return make([]int, 0)
	}

	input := string(bs)

	// do stuff with the input
	intInstr := make([]int, len(input))

	for i, val := range input {
		intInstr[i] = int(val) - 48
	}

	return intInstr
}

func countDigit(layer [6][25]int, digit int) int {
	digits := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			if layer[i][j] == digit {
				digits++
			}
		}
	}
	return digits
}

func buildLayers(instr []int) [][6][25]int {
	image := make([][6][25]int, 0)
	exit := false
	instrPointer := 0
	instrEnd := len(instr)

	// build layers
	for exit != true {
		var layer [6][25]int
		for i := 0; i < 6; i++ {
			for j := 0; j < 25; j++ {
				if instrPointer >= instrEnd {
					exit = true
					break
				} else {
					layer[i][j] = instr[instrPointer]
					instrPointer++
				}
			}
		}
		if instrPointer >= instrEnd {
			exit = true
			break
		} else {
			image = append(image, layer)
		}
	}
	return image
}

func part1(image [][6][25]int) {
	leastZerosLayer := 0
	layerZeros := make([]int, len(image))

	for i, layer := range image {
		layerZeros[i] = countDigit(layer, 0)
		if layerZeros[i] < layerZeros[leastZerosLayer] {
			leastZerosLayer = i
		}
	}

	ones := countDigit(image[leastZerosLayer], 1)
	twos := countDigit(image[leastZerosLayer], 2)
	product := ones * twos

	fmt.Println(layerZeros)
	fmt.Println(image[leastZerosLayer], ones, twos)

	fmt.Printf("Layer with the least zeros is %v and a result of %v\n", leastZerosLayer, product)
}
