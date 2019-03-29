package main

import (
    "fmt"
    "encoding/csv"
    "io"
    "log"
    "os"
    "time"
)

type Activity struct {
    name string
    duration time.Duration
    note string
}

func parseRecord(r []string) Activity {
    var act Activity
    act.name = r[0]
    d := fmt.Sprintf("%sh%sm", r[1][0:2], r[1][3:])
    act.duration, _ = time.ParseDuration(d)
    act.note = r[4]
    return act
}

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
        if len(record) < 1 {
            break
        }
        act := parseRecord(record)
        fmt.Printf("%s\t%s\t%v\t%s\n", act.name, record[1], act.duration, act.note)
    }
}
