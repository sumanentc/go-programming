package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	s := []int{9, 3, 5, 6, 7, 8, 2}
	fmt.Println("Befor Sort :: ", s)
	sort.Ints(s)
	fmt.Println("After Sort :: ", s)

	p1 := Person{
		"James",
		32,
	}

	p2 := Person{
		"Billy",
		25,
	}

	people := []Person{p1, p2}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
