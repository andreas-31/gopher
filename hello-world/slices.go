package main

import "fmt"

func printSlice( s []int) {
    fmt.Printf("%p - len: %d cap: %d %#v\n",
        s, len(s), cap(s), s)
}

// slice hat Typ [][]int
func printSliceNew(name string, slice ...[]int) {
    fmt.Println(name, ":")
    for _, s := range slice {
        fmt.Printf("%p - len: %d cap: %d - %#v\n", s, len(s), cap(s), s)
    }
}

func main() {
    var a []int
    printSlice(a)
    a = append(a, 1)
    printSlice(a)
    a = append(a, 2)
    printSlice(a)
    a = append(a, 3)
    printSlice(a)
    
    a2 := []int{}
    printSlice(a2)
    a2 = append(a2, 0, 1, 2) // [0 1 2]
    printSlice(a2)
    b2 := []int{3, 4, 5}
    a2 = append(a2, b2...) // append als variadische Funktion
    printSlice(a2)
    
    // Definition as composite literal
    b3 := []int{}
    printSlice(b3)
    c3 := []int{1, 2, 3}
    printSlice(c3)
    
    // Definition of slice with make
    d := make([]int, 2, 8)
    printSlice(d)
    
    for _, v := range d {
        fmt.Println(v)
    }
    
    // Slice a slice
    fmt.Println("\nSlice a slice")
    a4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    b4 := a4[1:5] // Slices a4 und b4 zeigen auf das gleiche Array
    printSliceNew("b4 := a4[1:5]", a4, b4) // Pointer a4 und b4 zeigen auf
                                           // verschiedene Elemente im Array
                // nämlich auf das erste Element des jeweiligen Slices
    b4 = append(b4, 11)
    printSliceNew("nach append", a4, b4)
    c4 := a4[1:5:5]
    printSliceNew("c4 = a4[1:5:5]", c4)
    c4 = append(c4, 14)
    printSliceNew("nach c4 append", a4, c4)
}

