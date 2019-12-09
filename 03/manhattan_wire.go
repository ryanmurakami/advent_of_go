package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input"
	bs, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(bs)
	lines := strings.Split(input, "\n")
	coordsA := strings.Split(lines[0], ",")
	coordsB := strings.Split(lines[1], ",")

	vectorsA := buildVectors(coordsA)
	vectorsB := buildVectors(coordsB)

	stepVectorA := make([]int, len(coordsA))
	stepVectorB := make([]int, len(coordsB))

	start := []int{0, 0}

	for i := 0; i < len(vectorsA); i++ {
		if i == 0 {
			stepVectorA[i] = getSteps(coordsA[i])
		} else {
			stepVectorA[i] = stepVectorA[i-1] + getSteps(coordsA[i])
		}
		stop := navigate(coordsA[i], start)
		vectorsA[i] = [][]int{[]int{start[0], start[1]}, []int{stop[0], stop[1]}}
		start = stop
	}

	start = []int{0, 0}

	for i := 0; i < len(vectorsB); i++ {
		if i == 0 {
			stepVectorB[i] = getSteps(coordsB[i])
		} else {
			stepVectorB[i] = stepVectorB[i-1] + getSteps(coordsB[i])
		}
		stop := navigate(coordsB[i], start)
		vectorsB[i] = [][]int{[]int{start[0], start[1]}, []int{stop[0], stop[1]}}
		start = stop
	}

	fmt.Println(vectorsA)
	fmt.Println(vectorsB)

	// matches := make([][2]int, 4)
	matches := make([]int, 4)

	for i := 0; i < len(vectorsA); i++ {
		for j := 0; j < len(vectorsB); j++ {
			crossing, aStepsDiff, bStepsDiff := areCrossing(vectorsA[i], vectorsB[j])
			if crossing == true {
				// matches = append(matches, match)
				var aSteps int
				var bSteps int
				if i == 0 {
					aSteps = 0
				} else {
					aSteps = stepVectorA[i-1]
				}
				if j == 0 {
					bSteps = 0
				} else {
					bSteps = stepVectorB[j-1]
				}
				matches = append(matches, aSteps+bSteps+aStepsDiff+bStepsDiff)
			}
		}
	}

	fmt.Println(matches)

	closestMatch := 0

	for _, match := range matches {
		// dist := calculateMhtn(match)
		if closestMatch == 0 || match < closestMatch {
			closestMatch = match
		}
	}

	fmt.Println(closestMatch)
}

func buildVectors(coords []string) [][][]int {
	vects := make([][][]int, len(coords))
	return vects
}

func navigate(direction string, currLoc []int) []int {
	dir := direction[0:1]
	nextLoc := []int{currLoc[0], currLoc[1]}
	distance := direction[1:len(direction)]
	distInt, err := strconv.Atoi(distance)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Compare(dir, "U") == 0 {
		nextLoc[1] = nextLoc[1] + distInt
	} else if strings.Compare(dir, "D") == 0 {
		nextLoc[1] = nextLoc[1] - distInt
	} else if strings.Compare(dir, "L") == 0 {
		nextLoc[0] = nextLoc[0] - distInt
	} else if strings.Compare(dir, "R") == 0 {
		nextLoc[0] = nextLoc[0] + distInt
	}

	return nextLoc
}

func getSteps(direction string) int {
	distance := direction[1:len(direction)]
	dist, _ := strconv.Atoi(distance)
	return dist
}

// crossing, aStepsDiff, bStepsDiff
func areCrossing(coordsA [][]int, coordsB [][]int) (bool, int, int) {
	aStepsDiff := 0
	bStepsDiff := 0
	matched := false
	vertA := isVert(coordsA[0], coordsA[1])
	vertB := isVert(coordsB[0], coordsB[1])

	if (vertA && vertB) || (!vertA && !vertB) {
		// going the same direction
		return false, 0, 0
	} else if vertA {
		if crossing(coordsA, coordsB) {
			aStepsDiff = intAbs(intAbs(coordsA[0][1]) - intAbs(coordsB[0][1]))
			bStepsDiff = intAbs(intAbs(coordsB[0][0]) - intAbs(coordsA[0][0]))
			matched = true
		}
	} else {
		if crossing(coordsB, coordsA) {
			aStepsDiff = intAbs(intAbs(coordsA[0][0]) - intAbs(coordsB[0][0]))
			bStepsDiff = intAbs(intAbs(coordsB[0][1]) - intAbs(coordsA[0][1]))
			matched = true
		}
	}

	return matched, aStepsDiff, bStepsDiff
}

func isVert(start []int, stop []int) bool {
	if start[1] == stop[1] {
		return false
	}
	return true
}

// a is vertical
func crossing(a [][]int, b [][]int) bool {
	yCross := ((a[0][1] < b[0][1]) && (b[0][1] < a[1][1])) ||
		((a[0][1] > b[0][1]) && (b[0][1] > a[1][1]))
	xCross := ((b[0][0] < a[0][0]) && (b[1][0] > a[0][0])) ||
		((b[0][0] > a[0][0]) && (b[1][0] < a[0][0]))
	return yCross && xCross
}

func calculateMhtn(match [2]int) int {
	a := math.Abs(float64(match[0]))
	b := math.Abs(float64(match[1]))
	sum := int(a) + int(b)
	return sum
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
