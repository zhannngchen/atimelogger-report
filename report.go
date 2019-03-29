package main

import (
    "fmt"
    "encoding/csv"
    "io"
    "log"
    "os"
)

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    r := csv.NewReader(f)
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(record)
    }
}
