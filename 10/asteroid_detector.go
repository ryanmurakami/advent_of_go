package main

// a lot of help came from reading rbusquet's code
// https://github.com/rbusquet/advent-of-code/tree/d907cb218f22b6b35f7d5ae4b8da020fba516f16/2019

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	astMap := readInput("input")

	mostVisibleAsteroids := 0
	bestX := 0
	bestY := 0

	for y, slice := range astMap {
		for x, spot := range slice {
			if spot == []rune("#")[0] {
				viz, _ := countVisibleAsteroids(x, y, astMap)
				if viz > mostVisibleAsteroids {
					mostVisibleAsteroids = viz
					bestX = x
					bestY = y
				}
			}
		}
	}

	fmt.Printf("Max asteroids visible %v at %v, %v\n", mostVisibleAsteroids, bestX, bestY)

	theX, theY := eliminate(bestX, bestY, astMap)
	fmt.Printf("Eliminated the 200th at %v, %v\n", theX, theY)
}

func readInput(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return make([]string, 0)
	}

	input := string(bs)
	astMap := strings.Split(input, "\n")

	return astMap
}

func countVisibleAsteroids(x int, y int, astMap []string) (int, map[float64]int) {
	seenAsts := make(map[float64]int)

	// starting clockwise
	for i := 0; i < len(astMap[0])-x; i++ {
		for j := 0; j >= -y; j-- {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				seenAsts[angle] = 1
			}
		}
	}

	for i := 0; i < len(astMap[0])-x; i++ {
		for j := 0; j < len(astMap)-y; j++ {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				seenAsts[angle] = 1
			}
		}
	}

	for i := 0; i >= -x; i-- {
		for j := 0; j < len(astMap)-y; j++ {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				seenAsts[angle] = 1
			}
		}
	}

	for i := 0; i >= -x; i-- {
		for j := 0; j >= -y; j-- {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				seenAsts[angle] = 1
			}
		}
	}

	return len(seenAsts), seenAsts
}

func isVisible(x int, y int, x2 int, y2 int, astMap []string) (bool, float64) {
	if x2 == x && y2 == y {
		return false, 0
	} else if string(astMap[y2][x2]) == "#" {
		angle := math.Atan2(float64(x2-x), float64(y2-y))
		return true, angle
	}
	return false, 0
}

func eliminate(x int, y int, astMap []string) (int, int) {
	observer := Asteroid{x, y}
	allAsts := make(map[float64][]Asteroid)

	// capture all asts
	for i := 0; i < len(astMap[0])-x; i++ {
		for j := 0; j >= -y; j-- {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				if allAsts[angle] == nil {
					allAsts[angle] = []Asteroid{}
				}
				allAsts[angle] = append(allAsts[angle], Asteroid{x + i, y + j})
			}
		}
	}

	for j := 0; j < len(astMap)-y; j++ {
		for i := 0; i < len(astMap[0])-x; i++ {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				if allAsts[angle] == nil {
					allAsts[angle] = []Asteroid{}
				}
				allAsts[angle] = append(allAsts[angle], Asteroid{x + i, y + j})
			}
		}
	}

	for i := 0; i >= -x; i-- {
		for j := 0; j < len(astMap)-y; j++ {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				if allAsts[angle] == nil {
					allAsts[angle] = []Asteroid{}
				}
				allAsts[angle] = append(allAsts[angle], Asteroid{x + i, y + j})
			}
		}
	}

	for j := 0; j >= -y; j-- {
		for i := 0; i >= -x; i-- {
			visible, angle := isVisible(x, y, x+i, y+j, astMap)
			if visible == true {
				if allAsts[angle] == nil {
					allAsts[angle] = []Asteroid{}
				}
				allAsts[angle] = append(allAsts[angle], Asteroid{x + i, y + j})
			}
		}
	}

	rightHemAngles := []float64{}
	leftHemAngles := []float64{}

	for k := range allAsts {
		if k >= 0 {
			rightHemAngles = append(rightHemAngles, k)
		} else {
			leftHemAngles = append(leftHemAngles, k)
		}
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(rightHemAngles)))
	sort.Sort(sort.Reverse(sort.Float64Slice(leftHemAngles)))

	zapped := 0

	for {
		for _, angle := range rightHemAngles {
			if len(allAsts[angle]) > 0 {
				astI := findClosestAsteroid(observer, allAsts[angle])
				zapped++

				if zapped == 200 {
					return allAsts[angle][astI].x, allAsts[angle][astI].y
				}

				// remove it
				others := allAsts[angle]
				others[astI] = others[len(others)-1]
				allAsts[angle] = others[:len(others)-1]
			}
		}

		for _, angle := range leftHemAngles {
			if len(allAsts[angle]) > 0 {
				astI := findClosestAsteroid(observer, allAsts[angle])
				zapped++

				if zapped == 200 {
					return allAsts[angle][astI].x, allAsts[angle][astI].y
				}

				// remove it
				others := allAsts[angle]
				others[astI] = others[len(others)-1]
				allAsts[angle] = others[:len(others)-1]
			}
		}
	}
}

// Asteroid ast
type Asteroid struct {
	x int
	y int
}

func findDistance(a Asteroid, b Asteroid) float64 {
	diffX := float64(b.x - a.x)
	diffY := float64(b.y - a.y)
	return math.Sqrt(math.Pow(diffX, 2) + math.Pow(diffY, 2))
}

func findClosestAsteroid(mainAst Asteroid, others []Asteroid) int {
	closest := 0
	closestDist := float64(99999)

	for i, ast := range others {
		dist := findDistance(mainAst, ast)
		if dist < closestDist {
			closestDist = dist
			closest = i
		}
	}

	return closest
}
