package main

import (
    "log"
    "os"
    "io"
    "crypto/md5"
    "flag"
)

var (
    flagFile = flag.String("file", "", "Wenn gesetzt wird direkt dieses File verwendet")
    flagURL = flag.String("url", "", "Wenn gesetzt wird von dieser URL geladen")
) // flagFile and flagURL are of type *string


func main() {
    flag.Parse()
    var input io.Reader
    input = os.Stdin
    printMD5(input)
}

func printMD5(r io.Reader) {
    h := md5.New()
    _, err := io.Copy(h, os.Stdin)
    if err != nil {
        log.Println("Fehler beim Einlesen;", err)
        return
    }
    log.Printf("%x\n", h.Sum(nil))
}

