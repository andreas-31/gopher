package main

import (
    "log"
    "os"
    "io"
    "crypto/md5"
)

func main() {
    h := md5.New()
    _, err := io.Copy(h, os.Stdin)
    if err != nil {
        log.Println("Fehler beim Einlesen;", err)
        os.Exit(1)
    }
    log.Printf("%x\n", h.Sum(nil))
}
