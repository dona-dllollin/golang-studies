package main 

import (
	"fmt"
)

func Sum[V int | float32 | float64](numbers []V) V {
	var total V 
	for _, e := range numbers {
		total += e
	}

	return total
}

func SumNumbers1(m map[string]int64) int64 {
	var s int64 
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers2[k comparable, v int64 | float64](m map[k]v) v {
	var s v
	for _, v := range m {
		s += v
	}

	return s
}

 type Number interface {
 	int64 | float64
 }
 func SumNumbers3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
	s += v
	}
	return s
 } 

 type UserModel[T int | float64] struct {
 Name string
 Scores []T
 }
 func (m *UserModel[int]) SetScoresA(scores []int) {
  m.Scores = scores
 }
 func (m *UserModel[float64]) SetScoresB(scores []float64) {
  m.Scores = scores
 }



func main() {
	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	 fmt.Printf("Generic Sums with Constraint: %v and %v\n",
	 SumNumbers2(ints),
	 SumNumbers2(floats))

	 var m1 UserModel[int]
	 m1.Name = "Noval"
	 m1.Scores = []int{1, 2, 3}
	 fmt.Println("scores:", m1.Scores)
	 var m2 UserModel[float64]
	 m2.Name = "Noval"
	 m2.SetScoresB([]float64{10, 11})
	 fmt.Println("scores:", m2.Scores)

}


