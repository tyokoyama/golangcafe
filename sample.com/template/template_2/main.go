package main

import (
	"os"
	"text/template"
)

type Member struct {
	No int
	Name string
	Tech string
}

// プログラム内で定義したtemplateを使って出力
func main() {
	// rangeキーワードを使うと繰り返し処理ができる。
	const template_text = `
No, 名前, 好きなGoogle技術,
{{range $index, $element := .}}{{$index}}, {{$element.Name}}, {{$element.Tech}},
{{end}}
`

	members := make([]Member, 3)
	members[0] = Member{No:1, Name: "ttyokoyama", Tech: "Appengine"}
	members[1] = Member{No:2, Name: "taknb2nch", Tech: "Go"}
	members[2] = Member{No:3, Name: "qt_luigi", Tech: "Go"}

	t := template.Must(template.New("sample").Parse(template_text))

	if err := t.Execute(os.Stdout, members); err != nil {
		println("t.Execute Error: ", err)
	}

}