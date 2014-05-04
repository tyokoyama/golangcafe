package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	//cmd := exec.Command("cmd", "/c", "echo hello")
	cmd := exec.Command("/bin/bash")
	//cmd := exec.Command("cmd")

	stdout, _ := cmd.StdoutPipe()
	stdin, _ := cmd.StdinPipe()

	cmd.Start()

	r := bufio.NewReader(stdout)
	w := bufio.NewWriter(stdin)

	ch := make(chan string)

	func1(r, ch)

	go func() {
		w.Write([]byte("ls\n"))
		w.Flush()
		stdin.Close()
	}()
	fmt.Println(<-ch)

	time.Sleep(5 * time.Second)

	fmt.Println("cmd2")
	func1(r, ch)
	go func() {
		w.Write([]byte("echo 'Hello World'\n"))
		w.Flush()
		stdin.Close()
	}()

	fmt.Println(<-ch)

	cmd.Wait()
}

func func1(r *bufio.Reader, ch chan<- string) {
	go func() {
		buf := make([]byte, 0)

		for {
			line, n, err := r.ReadLine()

			//fmt.Println(string(line))
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else if err != nil {
				fmt.Println(".", n, err)
				break
			}

			buf = append(buf, line...)
		}

		ch <- string(buf)

		fmt.Println("###############")
	}()
}