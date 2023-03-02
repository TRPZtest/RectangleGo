package main

import (
	"fmt"
	"math"
	"math/rand"
)

type point struct {
	x int64
	y int64
}

type vector struct {
	x, y int64
}

type diagonal struct {
	a, b point
}

func isRightAngle(a, b, c point) bool {
	ac := vector{c.x - a.x, c.y - a.y}
	bc := vector{c.x - b.x, c.y - b.y}

	dotProd := float64(ac.x*bc.x + ac.y*bc.y) //dot product of vectors

	modulesProd := math.Sqrt(float64(ac.x*ac.x+ac.y*ac.y)) * math.Sqrt(float64(bc.x*bc.x+bc.y*bc.y)) //prod of modules

	cos := dotProd / modulesProd

	return isEqual(cos, 0.0)
}

func isEqual(x, y float64) bool {
	e := 1e-3 //error
	if x == y {
		return true
	}
	d := math.Abs(x - y)

	result := d < e

	return result
}

//this function checks equals of pposite sides and then checks one angle
func isRectangle(diagonalX, diagonalY diagonal) bool {

	sideA := getDistanceBetweenPoints(diagonalX.a, diagonalY.a)
	sideC := getDistanceBetweenPoints(diagonalX.b, diagonalY.b)

	if !isEqual(sideA, sideC) {
		return false
	}

	sideB := getDistanceBetweenPoints(diagonalX.a, diagonalY.b)
	sideD := getDistanceBetweenPoints(diagonalX.b, diagonalY.a)

	if !isEqual(sideB, sideD) {
		return false
	}

	if !isRightAngle(diagonalX.a, diagonalX.b, diagonalY.a) {
		return false
	}

	return true
}

func getDistanceBetweenPoints(a, b point) float64 {
	d := math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2)

	return math.Sqrt(d)
}

func shufflePoints(points []point, rand rand.Rand) {
	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})
}

func remove(slice []point, n int) []point {

	slice = slice[:len(slice)-n]
	return slice
}

// func []quadrilateral getAllPossibleQuadrilateral

func main() {

	seed := 255

	rand := rand.New(rand.NewSource(int64(seed)))

	points := []point{{-5, 1}, {-1, 5}, {4, 0}, {0, -4}, {2, 2}, {0, 0}, {1, 0}, {1, 1}, {0, 1}, {2, 0}, {2, 1}, {11, 23}}

	shufflePoints(points, *rand)

	points = remove(points, rand.Intn(len(points)))

	diagonals := []diagonal{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			diagonalTemp := diagonal{points[i], points[j]}
			diagonals = append(diagonals, diagonalTemp)
		}
	}

	var count int

	for i := 0; i < len(diagonals); i++ {
		for j := i + 1; j < len(diagonals); j++ {
			if isRectangle(diagonals[i], diagonals[j]) {
				count++
			}
		}
	}

	fmt.Println("Number of rectangles:")
	fmt.Println(count)
}
