package main

import (
	"os"
	"os/exec"
	"io"
	"log"
	"time"
	"fmt"
)

func main() {
	cmd := exec.Command("cat")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func(in io.WriteCloser) {
		for i := 0; i < 5; i++ {
			in.Write([]byte(fmt.Sprintf("hoge_%d\n", i)))
			time.Sleep(1 * time.Second)
		}
		in.Close()
	}(stdin)

	go io.Copy(os.Stdout, stdout)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}