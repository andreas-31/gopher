package main

import (
    "bytes"
    "strings"
    "testing"
)

func TestPrintMD5(t *testing.T) {
    in := strings.NewReader("123")
    buf := &bytes.Buffer{}
    printMD5(in, buf)
    want := "202cb962ac59075b964b07152d234b70\n" // echo -n "123" | md5sum -
    got := buf.String()
    if got != want {
        t.Errorf("printMD5()=%s\nwant:%s\n", got, want)
    }
}
