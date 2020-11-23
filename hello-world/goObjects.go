package main

import "fmt"
import "https://github.com/andreas-31/gopher/hello-world/bar"

type rechteck struct {
    laenge int
    breite int
}

func flaeche(r rechteck) int {
    return r.laenge * r.breite
}

func (r rechteck) flaeche() int {
    return r.laenge * r.breite
}

func (r *rechteck) setBreite(b int) {
    r.breite = b
}

func main() {
    r := rechteck{laenge: 10, breite: 5}
    fmt.Println(r.flaeche())
    // Output: 50
    r.setBreite(100)
    fmt.Println(r.flaeche())
    
    // call func from imported package
    bar.Bar()
}
