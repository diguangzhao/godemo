package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func SortPeople(people []Person) {
	sort.Sort(ByAge(people))
}

func main()  {
	people := []Person{{"a", 1}, {"b", 2}, {"c", 4}, {"d", 3}}
	fmt.Println(people)
	SortPeople(people)
	fmt.Println(people)

	ints := []int{1, 4, 2, 3, 5}
	fmt.Println(ints)
	SortInt(ints)
	fmt.Println(ints)
}

func SortInt(ints []int)  {
	sort.Ints(ints)
}