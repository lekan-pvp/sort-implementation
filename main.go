package main

import (
	"fmt"

	"github.com/lekan/sort-implementation/mysort"
)

// sorting of int type
func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := mysort.IntSlice(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is %v\n", a)
}

//sorting of string type
func strings() {
	data := []string{"Monday", "Friday", "Tuesday", "Wednesday", "Sunday", "Thursday", "", "Saturday"}
	a := mysort.StringSlice(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is %v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int {
	return len(p.data)
}

func (p *dayArray) Less(i, j int) bool {
	return p.data[i].num < p.data[j].num
}

func (p *dayArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

// sorting of custom type day
func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	mysort.Sort(&a)
	if !mysort.IsSorted(&a) {
		panic("fail")
	}
	for _, i := range data {
		fmt.Printf("%s ", i.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()
}
