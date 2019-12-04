package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)
	lines := strings.Fields(str)
	total := 0

	for _, val := range lines {
		x, _ := strconv.Atoi(val)
		totalfuel := 0
		tempfuel := calcfuel(x)
		loop := true
		for loop {
			if tempfuel > 0 {
				totalfuel += tempfuel
				tempfuel = calcfuel(tempfuel)
			} else {
				loop = false
			}
		}
		total += totalfuel
	}

	fmt.Println(total)
}

func calcfuel(mass int) int {
	xdiv3 := mass / 3
	xdiv3float := float64(xdiv3)
	xdown := math.Floor(xdiv3float)
	finalval := int(xdown) - 2
	return finalval
}
