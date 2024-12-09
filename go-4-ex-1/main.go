package main


import (
	"fmt"
	"math"
)

// TODO: implement the function computeGrade
func computeGrade(gotPoints, maxPoints float64) float64 {
	if maxPoints == 0 {
		return 1.0
	}

	grade := 1.0 + 5.0*(gotPoints/maxPoints)
	return math.Min(math.Max(grade, 1.0), 6.0)
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 20.0)) 
	fmt.Println(computeGrade(12.0, 20.0)) 
	fmt.Println(computeGrade(3.0, 28.0)) 
}
