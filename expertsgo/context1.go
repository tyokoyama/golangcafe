package main

import (
	"context"
	"fmt"
)

func main() {
	empty := context.Background()
	
	cancelCtx, cancel := context.WithCancel(emptyCtx)
	defer cancel()
	doSomeThing(cancelCtx)
}