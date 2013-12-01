package main

type Name string

func (n Name) Print() {
	println(n)
}

func main() {
	str :=  "Hello" +
	" World!" +
	" Hoge!"

	n := Name("hoge")

	n.
		Print()

	println(str)

	a := 0

	b := (a++)

	println(b)
}