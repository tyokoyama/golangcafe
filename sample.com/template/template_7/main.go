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

	t := template.Must(template.ParseFiles("sample.txt", "sample2.txt", "sample3.txt"))

	// if err := t.Execute(os.Stdout, members); err != nil {
	// 	println("t.Execute Error: ", err)
	// }

	if err := t.ExecuteTemplate(os.Stdout, "sample2.txt", members); err != nil {
		println("t.Execute Error: ", err)
	}

}