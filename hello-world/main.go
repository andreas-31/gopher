// create go.mod
// go mod init GIT_CLONE_URL
// go mod init local/meinModul

// create main.go
package main

import "fmt"
import "os"

func main() {
    fmt.Println("Hello, World!")
    
    // print it!
    fmt.Print("Hallo ", "Print()\n")
    fmt.Println("Hallo Println()")
    var s = "Printf()"
    fmt.Printf("Hallo, %s\n", s)
    var nr = 2
    var name = "Till"
    fmt.Printf("%03d: Hallo, %s\n", nr, name)
    // Wichtige Verben: 
    // %v: jeder Variablentyp
    // %+v: bei Strukturen auch Feldnamen ausgeben
    // %#v: Ausgabe als Go-Code
    // %T: gibt Typ der Variablen aus
    // %b: Ausgabe als binäre Zahl
    // %d: Ausgabe zur Basis 10
    // %03d oder %03b: Zahlen mit 0 auffüllen
    // %p: Adresse der Variablen im Speicher (Pointer)
    // %s: direkte Ausgabe eines Strings
    // %x: Ausgabe String als hexadezimaler Wert
    // Weitere: https://golang.org/pkg/fmt/
    
    // String Format (print string): fmt.Sprint...
    nr = 5
    s = fmt.Sprintf("%03d: Hallo, %s (Sprint)\n", nr, name)
    fmt.Print(s)
    
    // Print into file: fmt.Fprint...
    file, _ := os.Create("datei.txt") // oder os.Stdout
    fmt.Fprintf(file, "%03d: Hallo, %s\n", nr, name)
    file.Close()
    
    // Variablen werden auf Deault-Werte initialisiert
    //var nummer int // 0
    //var txt string // ""
    //var checked bool // false
    //var meinUser *user // nil - Pointer
    //var liste []string // nil - Slice
    
    // Definition einer Gruppe von Variablen
    /*var (
        home = os.Getenv("HOME")
        user = os.Getenv("USER")
        gopath = os.Getenv("GOPATH")
    )*/
    
    // Kurdeklaration von Variablen
    //nummer := 10 // int
    //name := getName() // string
    //a, b := 12, 34
    //a, b = b, a // Werte tauschen
    
    
}

// Programm gofmt
// Programm goimports
// go build
// go run main.go
