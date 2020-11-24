package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    // wurde mind. 1 argument übergeben?
    if len(os.Args) == 1 {
        fmt.Println("Mindestens eine Datei als Parameter erwartet")
        os.Exit(1)
    }
    
    for _, fname := range os.Args[1:] {
        fmt.Println(fname)
    }
    
    // Datei öffnen
    fname := os.Args[1]
    f, err := os.Open(fname) // func Open(name string) (*File, error)
    if err != nil {
        log.Printf("Fehler beim Öffnen der Datei: %s\n%s\n", fname, err)
        f.Close()
        //continue
    }
}
