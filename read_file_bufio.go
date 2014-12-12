package main

import "bufio"
import "fmt"
import "os"
import "io"

func main() {
    f, _ := os.Open("sample.log")
    defer f.Close()
    buf_rdr := bufio.NewReader(f)
    line, err := buf_rdr.ReadString('\n')
    for err == nil {
        line, err = buf_rdr.ReadString('\n')
        fmt.Println(line)
    }
    if err == io.EOF {
        fmt.Println(line)
    }
    if err != nil {
        fmt.Println(err)
    }
}