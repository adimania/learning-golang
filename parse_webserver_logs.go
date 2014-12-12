// sample.log at http://pastie.org/pastes/7752195/text?key=rfuehqhwmrxhiyrafpislg
package main

import "bufio"
import "fmt"
import "os"
import "strings"

func main() {
    f, _ := os.Open("sample.log")
    defer f.Close()
    ret_codes := make(map[string]int)
    buf_rdr := bufio.NewReader(f)
    line, err := buf_rdr.ReadString('\n')
    for err == nil {
        if line != "" {
            words := strings.Split(line, " ")
            value, status := ret_codes[words[9]]
            if status {
                ret_codes[words[9]] = value + 1
            } else {
                ret_codes[words[9]] = 1
            }
        }
        line, err = buf_rdr.ReadString('\n')
    }
    fmt.Print(ret_codes)
}