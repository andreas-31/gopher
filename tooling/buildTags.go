// +build windows
package myPackage

func ipCmd() {
    // ...
    out, err := exec.Command("ipconfig", "/all").Output()
    // ...
}
