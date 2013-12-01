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

func (m Member) PrintFlag() (result string) {
	if m.Flag {
		result = "*,"
	} else {
		result = ""
	}

	return
}

func main() {
	members := make([]Member, 3)
	members[0] = Member{No:1, Name: "ttyokoyama", Tech: "Appengine", Flag: false}
	members[1] = Member{No:2, Name: "taknb2nch", Tech: "Go", Flag: true}
	members[2] = Member{No:3, Name: "qt_luigi", Tech: "Go", Flag: false}

	t := template.Must(template.ParseFiles("sample.txt", "sample2.txt"))

	// templates := t.Templates()
	// for _, temp := range templates {
	// 	println(temp.Name())
	// }

	// テンプレートの名前はファイル名が付けられる。
	t2 := t.Lookup("sample2.txt")
	// if t2 != nil {
	// 	println(t2.Name())
	// } else {
	// 	println("nil")
	// }

	if err := t2.Execute(os.Stdout, members); err != nil {
		println("t.Execute Error: ", err)
	}

}