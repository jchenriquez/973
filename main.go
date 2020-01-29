package main

import "fmt"
import "math"
import "container/heap"

type PointsPriority [][]int

func (points PointsPriority) Len() int {
  return len(points)
}

func (points PointsPriority) Less(i, j int) bool {
  return getDistance(points[i], []int{0,0}) > getDistance(points[j], []int{0,0})
}

func (points PointsPriority) Swap(i, j int) {
  points[i], points[j] = points[j], points[i]
}

func (points *PointsPriority) Push(point interface{}) {
  *points = append(*points, point.([]int))
}

func (points *PointsPriority) Pop() interface {} {
  point := (*points)[len(*points)-1]
  *points = (*points)[:len(*points)-1]
  return point
}

func (points PointsPriority) Peek() interface{} {
  return points[0]
}

func getDistance(point []int, point2 []int) float64 {
  return math.Sqrt(math.Pow(float64(point[0]-point2[0]), 2)+math.Pow(float64(point[1] - point2[1]), 2))
}

func kClosest(points [][]int, K int) [][]int {
  result := make(PointsPriority, 0, K)
  pP := &result
  heap.Init(pP)
  centerPoint := []int{0,0}

  for _, point := range points {
    if len(*pP) == K {
      pointK := PointsPriority(result).Peek()

      if getDistance(pointK.([]int), centerPoint) > getDistance(point, centerPoint) {
        heap.Pop(pP)
        heap.Push(pP, point)
      }
    } else {
      heap.Push(pP, point)
    }
  }

  return [][]int(*pP)
}

func main() {
  fmt.Printf("kClosests: %v\n", kClosest([][]int{{3,3}, {5, -1}, {-2,4}, {3,2}}, 2))
}