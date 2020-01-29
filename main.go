package main

import "fmt"
import "sort"
import "math"

func getDistance(point []int, point2 []int) float64 {
  return math.Sqrt(math.Pow(float64(point[0]-point2[0]), 2)+math.Pow(float64(point[1] - point2[1]), 2))
}

func kClosest(points [][]int, K int) [][]int {
   pointsCpy := make([][]int, len(points))   
   copy(pointsCpy, points)

  sort.Slice(pointsCpy, func(i, j int)bool {
    return getDistance(pointsCpy[i], []int{0,0}) < getDistance(pointsCpy[j], []int{0,0})
  })

  return pointsCpy[:K]
}

func main() {
  fmt.Printf("kClosests: %v\n", kClosest([][]int{{3,3}, {5, -1}, {-2,4}, {3,2}}, 2))
}