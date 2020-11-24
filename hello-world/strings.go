package main

import "fmt"

func main() {
    helloWorld := "你好，世界" // string ist ein []byte in UTF-8 kodiert
    welcomeToChina := "欢迎来到中国"
    myStrings := []string{helloWorld, welcomeToChina}
    
    for _, s := range myStrings {
        fmt.Println(s)
        fmt.Println(len(s))
        for i, v := range s { // index zeigt immer auf das erste Byte des Zeichens
            fmt.Printf("Type of v: %T\n", v) // v ist von Typ (u?)int32 (rune)
            fmt.Println(i, v)
        }
    }
}
