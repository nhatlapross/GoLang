package main

import "fmt"

func main() {
	points := [3]int{7, 10, 5}
	fmt.Printf("%v,%T\n", points, points)

	points2 := [...]int{7, 10, 5, 9, 2}
	fmt.Printf("%v,%T\n", points2, points2)

	var point3 [3]string
	point3[0] = "Alex"
	point3[1] = "Tom"
	point3[2] = "Yuh"
	fmt.Printf("%v,%T\n", point3, point3)
	fmt.Printf("%v,%T\n", point3[2], point3[2])
	fmt.Println(len(point3))

	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades:%v", grades)

	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("length: %v\n", cap(a))

	a = append(a, 1)
}
