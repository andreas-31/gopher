package main

import (
    "fmt"
    "log"
    "os"
    "io"
    "crypto/md5"
    "flag"
    "net/http"
)

var (
    flagFile = flag.String("file", "", "Wenn gesetzt wird direkt dieses File verwendet")
    flagURL = flag.String("url", "", "Wenn gesetzt wird von dieser URL geladen")
) // flagFile and flagURL are of type *string


func main() {
    flag.Parse()
    var input io.Reader = os.Stdin
    var output io.Writer = os.Stdout
    
    switch {
        case *flagFile != "":
            fd, err := os.Open(*flagFile)
            if err != nil {
                fmt.Fprintln(output, "Fehler beim Öffnen des Files: ", err)
                os.Exit(1)
            }
            defer fd.Close()
            input = fd
        case *flagURL != "":
            resp, err := http.Get(*flagURL)
            if err != nil {
                fmt.Fprintln(output, "Fehler beim Laden: ", err)
                os.Exit(1)
            }
            defer resp.Body.Close()
            input = resp.Body
    }
    
    printMD5(input, output)
}

func printMD5(r io.Reader, w io.Writer) {
    h := md5.New()
    _, err := io.Copy(h, r)
    if err != nil {
        log.Println("Fehler beim Einlesen;", err)
        return
    }
    fmt.Fprintf(w, "%x\n", h.Sum(nil))
}

