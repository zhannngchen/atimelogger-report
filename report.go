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

var stats = make(map[string]map[string]time.Duration)

func parseRecord(r []string) Activity {
    var act Activity
    act.name = r[0]
    d := fmt.Sprintf("%sh%sm", r[1][0:2], r[1][3:])
    act.duration, _ = time.ParseDuration(d)
    act.note = r[4]
    return act
}

func updateStats(a Activity) {
    items := stats[a.name]
    if items == nil {
        items = make(map[string]time.Duration)
        stats[a.name] = items
    }
    items[a.note] += a.duration
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
            log.Print(err)
            break
        }
        if len(record) < 1 {
            break
        }
        act := parseRecord(record)
        updateStats(act)
    }

    for actName, item := range stats {
        fmt.Println(actName)
        for it, time := range item {
            fmt.Printf("\t%s\t%v\n", it, time)
        }
    }
}
