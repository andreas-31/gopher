package bar

import "fmt"

var Prefix = ""
var geheim = "meinGeheimnis"

func bar() {
    fmt.Println(Prefix, "bar")
}

func Bar() {
    bar()
}
