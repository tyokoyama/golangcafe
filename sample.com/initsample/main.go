package main

import (
	"fmt"
	_ "github.com/tyokoyama/golangcafe/sample.com/initsample/lib"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("main is called.")

	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(path)

	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}

}