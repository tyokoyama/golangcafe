package main

import (
	"fmt"
	"time"
	"errors"
	"runtime"
)

func doConcurrently(job string, err chan error) {
    go func() {
        fmt.Println("doing job", job)
        time.Sleep(1 * time.Second)
        err <- errors.New("something went wrong!")
    }()
}

func main() {
	runtime.GOMAXPROCS(2)

    jobs := []string{"one", "two", "three"}

    errc := make(chan error)
    for _, job := range jobs {
        doConcurrently(job, errc)
    }
    for _ = range jobs {
        if err := <-errc; err != nil {
            fmt.Println(err)
        }
    }
}