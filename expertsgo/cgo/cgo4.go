package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct {
	int age;
	const char *name;
} person_t;

person_t *new_person(int age, const char *name) {
	person_t *p = (person_t*)malloc(sizeof(person_t));
	p->age = age;
	p->name = name;
	return p;
}

void person_say(person_t *p) {
	fprintf(stderr, "Hello, I'm %s\n", p->name);
}
*/
import "C"

type Person struct {
	c *C.person_t
}

func NewPerson(age int, name string) *Person {
	c := C.new_person(C.int(age), C.CString(name))
	return &Person {c: c}
}

func (p *Person) Name() string {
	return C.GoString(p.c.name)
}
func (p *Person) SetName(name string) {
	p.c.name = C.CString(name)
}

func (p *Person) Say() {
	C.person_say(p.c)
}

func main() {
	p := NewPerson(10, "bob")
	p.SetName("alice")
	p.Say()
}