package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// 引数で指定したコマンドのパスを取得する。
	// インストールされていなければエラーが返る。
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	fmt.Printf("Path = %s\n", path)

	// ls -lを実行する。（対象のディレクトリは環境によって書き換えること。）
	cmd := exec.Command("ls", "-l", "/Users/yokoyama")

	// Pipeの取得は対象のプロセスの起動の前に終わらせておく。
	stdoutpipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("StdoutPipe Error: %v\n", err)
		return
	}
	defer stdoutpipe.Close()

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Command Start Error: %v\n", err)
		return
	}

	stdout, err := ioutil.ReadAll(stdoutpipe)
	if err != nil {
		fmt.Printf("Command Error: %v\n", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command Wait Error: %v\n", err)
		return
	}

	fmt.Printf("%s\n", stdout)
}