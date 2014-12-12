package main
import "fmt"
import "bufio"
import (
    "os"
    "io/ioutil"
)

func main() {
    data, _ := ioutil.ReadFile("sample.log")
    f, _ := os.Open("sample.log")
    scanner := bufio.NewScanner(f)
    scanner.Scan()
    fmt.Println(scanner.Text())
    f.Close()
}