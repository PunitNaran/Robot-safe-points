package main

import (
	"fmt"
	"strconv"
)

const (
	middlePoint = 720
)

func main() {
	xCord := makeRange(middlePoint*-1, middlePoint)
	yCord := make([]int, len(xCord)) // copy xCord because yCord is the same size with same values
	copy(yCord, xCord)
	safe := make([][]int, 0)
	for _, x := range xCord { // Iterate over x cordinates
		xSafe := make([]int, 0)
		sumX := sum(x)
		for _, y := range yCord { // iternate over y cordinates
			sumY := sum(y)
			digitSum := sumX + sumY // compute the x and y sum from digits
			if digitSum <= 23 {
				xSafe = append(xSafe, 1) // this is a potential safe spot
				continue
			}
			xSafe = append(xSafe, 0)
		}
		safe = append(safe, xSafe)
	}
	checkSafePoint(safe) // compute safe spots
}

func checkSafePoint(safe [][]int) {
	counter := 1 // first position is safe position
	safe[middlePoint][middlePoint] = 2
	for changing := true; changing; { // validate to safe points
		changing = false
		for y, xSafe := range safe {
			for x, valid := range xSafe {
				if y == 0 || x == 0 || y == len(safe)-1 || x == len(xSafe)-1 {
					continue // Ignore the padding
				}
				if safe[y][x] == 2 {
					continue // already a safe point
				}
				if valid == 1 && (safe[y+1][x] == 2 || safe[y][x+1] == 2 || safe[y-1][x] == 2 || safe[y][x-1] == 2) {
					safe[y][x] = 2 // se this as a safe point
					changing = true
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
}

func sum(num int) int {
	result := 0
	if num < 0 { // abs simple
		num *= -1
	}
	sNum := strconv.Itoa(num) // convert to string
	digits := make([]int, 0)
	for _, r := range sNum {
		digits = append(digits, int(r-'0')) // get int from rune
	}
	for _, v := range digits {
		result += v // add the digits
	}
	return result
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i // Create a range from a number of values
	}
	return a
}
