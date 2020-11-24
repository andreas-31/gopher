package main

import "fmt"

func intCopy(ziel []int, quelle []int) int {
    n := 0
    for i := range ziel {
        if i >= len(quelle) {
            break // oder gleich return n
        }
        ziel[i] = quelle[i]
        n++
    }
    return n
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    b := make([]int, 10)
    n := intCopy(b, a)
    fmt.Printf("%d Elemente kopiert: b=%#v\n", n, b)
}
