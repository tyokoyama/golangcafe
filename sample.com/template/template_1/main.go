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

func main() {

	// templateはTab、改行、空白全てが反映されるので注意
	const template_text = `今日の主催者は、{{.Name}}です。
`
	t := template.Must(template.New("sample").Parse(template_text))

	member := Member{No:1, Name: "ttyokoyama", Tech: "Appengine"}

	if err := t.Execute(os.Stdout, member); err != nil {
		println("t.Execute Error: ", err)
	}
}