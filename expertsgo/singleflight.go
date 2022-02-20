package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var g group

	g.do("foo")
	g.do("bar")
	g.Wait()

	g.Forget()
	g.do("hoge")
	g.do("fuga")
	g.Wait()
}

type group struct {
	singleflight.Group
	sync.WaitGroup
}

func (g *group) do(s string) {
	g.Add(1)

	go func() {
		defer g.Done()

		v, err, shared := g.Do("key", func() (interface{}, error) {
			time.Sleep(1 * time.Second)
			log.Printf("cached %s\n", s)
			return s, nil
		})

		log.Println(v, err, shared)
	}()
}

func (g *group) Forget() { g.Group.Forget("key")}