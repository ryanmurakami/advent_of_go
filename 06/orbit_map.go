package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(bs)
	instr := strings.Fields(input)

	orbitMap := make(map[string]string)

	// build orbitMap
	for _, orbit := range instr {
		planets := strings.Split(orbit, ")")
		orbitMap[planets[1]] = planets[0]
	}

	youPaths := makeOrbitChain(orbitMap, "YOU")
	santaPaths := makeOrbitChain(orbitMap, "SAN")

	youMatch := 0
	santaMatch := 0

	for i := 0; i < len(youPaths); i++ {
		for j := 0; j < len(santaPaths); j++ {
			if youPaths[i] == santaPaths[j] {
				fmt.Printf("found a match at %v %v\n", i, j)
				youMatch = i
				santaMatch = j
				break
			}
		}
		if youMatch > 0 {
			break
		}
	}

	fmt.Println(youMatch + santaMatch)

	// fmt.Printf("Total orbit count is %v\n", countOrbits(orbitMap))
}

func makeOrbitChain(orbitMap map[string]string, start string) []string {
	paths := make([]string, 0)

	target, hit := orbitMap[start]

	for hit == true {
		paths = append(paths, target)
		target, hit = orbitMap[target]
	}

	return paths
}

func countOrbits(orbitMap map[string]string) int {
	orbits := 0
	//process orbitMap

	for _, value := range orbitMap {
		target := value
		hit := true

		for hit == true {
			orbits++
			target, hit = orbitMap[target]
		}
	}

	return orbits
}
