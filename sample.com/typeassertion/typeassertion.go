package typeassertion

type Public interface {
	Name() string
}

type Example struct {
	name string
}

func (e Example) Nme() string {
//func (e Example) Name() string {
	return e.name
}

// func NewExample() Example {
// 	return Example{"No Name"}
// }

func NewExample() Public {
	return Example{"No Name"}
}

func NewExample2() Public {
	var p Public
	e := Example{"No Name"}
	p.(Example)
//	e.(Public)

	return e
}