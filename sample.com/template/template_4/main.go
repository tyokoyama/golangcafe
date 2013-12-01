package main

import (
	"os"
	"text/template"
)

type Member struct {
	No int
	Name string
	Tech string
	Flag bool
}

// プログラム内で定義したtemplateを使って出力
func main() {
	// lenキーワードは引数の長さを取得
	const template_text = `
計：{{len .}}件
No, 名前, 好きなGoogle技術,
{{range $index, $element := .}}{{$index}}, {{$element.Name}}, {{$element.Tech}},{{if .Flag}}*,{{end}}
{{end}}
`

	members := make([]Member, 3)
	members[0] = Member{No:1, Name: "ttyokoyama", Tech: "Appengine", Flag: false}
	members[1] = Member{No:2, Name: "taknb2nch", Tech: "Go", Flag: true}
	members[2] = Member{No:3, Name: "qt_luigi", Tech: "Go", Flag: false}

	t := template.Must(template.New("sample").Parse(template_text))

	if err := t.Execute(os.Stdout, members); err != nil {
		println("t.Execute Error: ", err)
	}

}