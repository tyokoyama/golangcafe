package main

import (
	"flag"
	"fmt"
)

func main() {
	// 以下の使い方はどれも同じ。先頭に"-"をつけた引数の意味を定義する。
	// flag.Bool
	// flag.Duration
	// flag.Float64
	// flag.Int
	// flag.Int64
	// flag.String
	// flag.Uint
	// flag.Uint64
	// 第1引数には"-"を取った名前を指定する。"-"をつけると"--x"という感じになる。
	// 第2引数は既定値（引数省略時など）
	// 第3引数は使い方がおかしい時（Parseできない時に表示するコメント）
	// 戻り値にコマンドパラメータの設定値が返される。
	// bool以外の型の引数の場合は-i hogeは、型が一致しないと、引数エラーになる。（次のパラメータを参照する）
	// コマンド例：go run flagsample.go -x -i=0 hoge
	//           go run flagsample.go -x=false hoge
	xflag := flag.Bool("x", true, "Usage: -x")
	iflag := flag.Int("i", 1, "Usage: -i")

	// 任意の変数に引数の値を設定する場合はBoolVar()を使う。他の型も存在する。
	var lflag bool
	flag.BoolVar(&lflag, "l", false, "Usage: -l")

	// 引数の解釈
	flag.Parse()

	// flag.Args()は引数の一覧を取得する。
	fmt.Println("-- Args -- ")
	fmt.Println(*xflag, *iflag, lflag, flag.Args())

	// フラグのデフォルト値を出力する。
	fmt.Println("-- PrintDefaults -- ")
	flag.PrintDefaults()

	// 指定された引数を全てチェック
	// VisitAll()を使うと、指定されていないものもすべてチェック
	fmt.Println("-- Visit -- ")
	flag.Visit(func(f *flag.Flag){
		fmt.Println(f)
	})
}