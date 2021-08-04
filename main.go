package main

import (
    "fmt"
    "strings"
)

var marks = []string{"|", "/", "-", "\\"}

func mark(i int) string {
    return marks[i%4]
}

func progress_bar(p int) string {
    bar := strings.Repeat("=", p / 5)

    if p != 100 {
        bar += ">"
    }

    return bar
}

func main() {
    cnt := 5000000000
    for i := 1; i <= cnt; i++ {
        if i%(cnt/100) == 0 {
            p := i / (cnt / 100)
            fmt.Printf("\rLoading: %s %2d%% [%-20s]", mark(p), p, progress_bar(p))
        }
    }
    fmt.Println("\nDone.")
}
