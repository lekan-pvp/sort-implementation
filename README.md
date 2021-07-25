# sort implementation

## mysort package

    // mysort package
    // mysort.go
    
    package mysort

    type Interface interface {
      Len() int
      Less(i, j int) bool
      Swap(i, j int)
    }

    func Sort(data Interface) {
      for pass := 1; pass < data.Len(); pass++ {
        for i := 0; i < data.Len()-pass; i++ {
          if data.Less(i+1, i) {
            data.Swap(i, i+1)
          }
        }
      }
    }

    func IsSorted(data Interface) bool {
      n := data.Len()
      for i := n - 1; i > 0; i-- {
        if data.Less(i, i-1) {
          return false
        }
      }
      return true
    }

    type IntSlice []int

    func (p IntSlice) Len() int {
      return len(p)
    }

    func (p IntSlice) Less(i, j int) bool {
      return p[i] < p[j]
    }

    func (p IntSlice) Swap(i, j int) {
      p[i], p[j] = p[j], p[i]
    }

    type StringSlice []string

    func (p StringSlice) Len() int {
      return len(p)
    }

    func (p StringSlice) Less(i, j int) bool {
      return p[i] < p[j]
    }

    func (p StringSlice) Swap(i, j int) {
      p[i], p[j] = p[j], p[i]
    }

## main.go

    //main.go
    
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
