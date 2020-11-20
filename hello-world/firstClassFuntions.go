package main

import "fmt"

func main() {
    f := func(s string) bool {
        if len(s) < 3 {
            return false
        }
        return true
    }
    
    s := []string{"a", "abcd", "abc", "ab"}
    fmt.Println(meinFilter(s, f))
    // Output: [abcd, abc]
    
    // nochmal mit Closure-Funktion
    f3 := laengenFilter2(3)
    f2 := laengenFilter(2)
    fmt.Println(meinFilter(s, f2))
    fmt.Println(meinFilter(s, f3))
    
    // Aufruf einer Funktion mit defer
    meineFunktion()
    meineFunktion2()
}

type filterFunc func(string)bool

// Implementation of function for filtering
func meinFilter(s []string, filter filterFunc) []string {
    var out []string
    for _, str := range s {
        if filter(str) {
            out = append(out, str)
        }
    }
    return out
}

// Closure
// Achtung: len() -> Anzahl der Bytes, nicht der Zeichen!
// besser: utf8.RuneCountInString()
func laengenFilter(laenge int) filterFunc {
    return func(s string) bool {
        if len(s) < laenge {
            return false
        }
        return true
    }
}
// alternative Signatur für Closure:
func laengenFilter2(laenge int) func(string) bool {
    return func(s string) bool {
        if len(s) < laenge {
            return false
        }
        return true
    }
}

// Funtion mit defer
func meineFunktion() {
    defer fmt.Println(1)
    defer fmt.Println(2)
    fmt.Println(3)
}
func meineFunktion2() {
    defer func() {
        fmt.Println(1)
        fmt.Println(2)
    }()
    fmt.Println(3)
    return
    fmt.Println(4)
}



