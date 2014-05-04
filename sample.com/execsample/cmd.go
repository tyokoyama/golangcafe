package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// 参考URL: http://goo.gl/tkRU1z
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
	// 下のコードだと、リアルタイムに出力されなくなる。
	// go func(out io.ReadCloser) {
	// 	buf := new(bytes.Buffer)

	// 	io.Copy(buf, out)

	// 	fmt.Println(buf.String())
	// }(stdout)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
