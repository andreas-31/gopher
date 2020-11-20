package main

import "fmt"
import "io"
import "errors"

func add2(a int) int {
    return 2 + a
}

func main() {
    const c = 2.0
    fmt.Println(add2(c))
    // Output:
    // 4
    var v = 2.0
    fmt.Println(add2(int(v)))
    // Output:
    // Fehler beim Kompilieren:
    // cannot use v (type float64) as type int in argument to add2
    // const a float64 = 2.0
    // const b = float64(2.0)
    
    // Zähler iota
    const (
        Montag = iota // 0
        Dienstag = iota // 1
        Mittwoch = iota // 2
        Donnerstag // 3
        Freitag // 4
        Samstag // 5
        Sonntag // 6
    )
    
    type ByteSize float64
    const (
        _ = iota // ignoriere 0
        KB ByteSize = 1 << (10 * iota) // 2^10 = 1024
        MB
        GB
        TB
        PB
        EB
        ZB
        YB
    )
    // skip 0, start left shift at 1
    // 0100 0000 0000
    //             |- 2^1 = 2
    //            |-- 2^2 = 4
    //           |--- 2^3 = 8
    //         |----- 2^4 = 16
    //        |------ 2^5 = 32
    //       |------- 2^6 = 64
    //      |-------- 2^7 = 128
    //    |---------- 2^8 = 256
    //   |----------- 2^9 = 512
    //  |------------ 2^10 = 1024
    
    // Pointer
    var a int
    var b *int // Typ Pointer auf int
    
    a = 123
    b = &a // Speicheradresse mit &
    fmt.Println(b, *b)
    *b = 100 // Dereferenzierung
    fmt.Println(a)
    
    a2 := foo()
    fmt.Println(a2, *a2)
    
    // Strukturen
    type adresse struct {
        strasse string
        stadt string
    }
    // Deklaration mit Feldnamen
    a3 := adresse {
        strasse: "Musterstr.",
        stadt: "Musterstadt",
    }
    // Deklaration mit allen Feldern
    b3 := adresse{"Bahnhofstr.", "Berlin"}
    fmt.Println(a3, b3)
    // Tags
    // Note: Typ und Felder sind groß geschrieben, daher beide public!
    type Adresse struct {
        Strasse string `json:"street"`
        Stadt string `json:"city"`
    }
    
    // Funktionen mit zwei Rückgabewerten
    fmt.Println(swap(11, 12))
    a4, b4 := swap(30, 31)
    fmt.Println(a4)
    fmt.Println(b4)
    
    // Aufruf einer Funktion mit benannten Rückgabewerten
    a5, b5 := ReturnSomething("Hallo Till!")
    fmt.Println(a5) // leer
    fmt.Println(b5) // Hallo Till!
    
    // Aufruf einer variadischen Funktion
    fmt.Println(addAll(1, 2))
    fmt.Println(addAll())
    fmt.Println(addAll(1, 2, 3, 4, 5))
    
}

// Funktion liefert einen Pointer zurück
func foo() *int {
    bar := 123
    return &bar
}

func add(a, b, c int) int {
    return a + b + c
}

func swap(a, b int) (int, int) {
    return b, a
}

// Funktion mit Fehlerbehandlung
func ReadSomething(r io.Reader) (int, error) {
    // ...
    n := 10
    someError := false
    if someError {
        return 0, errors.New("error when reading")
    }
    // ...
    return n, nil
}

// Funktion mit benannten Rückgabewerten
func ReturnSomething(s string) (a string, b string) {
    b = s // keine Deklaration notwendig/erlaubt
    return a, b // a = ""
}

type item struct {
    name string `json:"name"`
    city string `json:"city"`
}
// Funktion mit Naked Return
func findItem(items []item, name string) (result item, ok bool) {
    for _, i := range items {
        if i.name == name {
            result = i
            ok = true
            return
        }
    }
    return
}

func findItemBetterStyle(items []item, name string) (result item, ok bool) {
    for _, i := range items {
        if i.name == name {
            result = i
            return result, true
        }
    }
    return result, false
}

// Variadische Funktion
// Typ von a ist ein Slice: []int
func addAll(a ...int) int {
    summe := 0
    for _, v := range a {
        summe = summe + v
    }
    return summe
}



