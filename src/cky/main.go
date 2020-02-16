package main

import "fmt"

func main() {
	st := &St{
		Name: "link",
	}
	var i interface{} = st

	if o, ok := i.(*St); ok {
		fmt.Println("i is struct St")
		o.Ping()
		o.Pong()
		o.String()
	}

	if o, ok := i.(Inter); ok {
		fmt.Println("i is interface Inter ")
		o.Ping()
		o.Pong()
	}

	if o, ok := i.(Anter); ok {
		fmt.Println("i is interface Anter ")
		o.Ping()
		o.Pong()
		o.String()
	}
}

type Inter interface {
	Ping()
	Pong()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	fmt.Println("Ping")
}

func (St) Pong() {
	fmt.Println("Pong")
}

func (s *St) String() {
	fmt.Println(s.Name)
}

// --------------------- 方法集 --------------------------
type X struct {
	a int
}

type Z struct {
	*X
}

func (x X) Get() int {
	return x.a
}

func (x *X) Set(i int) {
	x.a = i
}

// --------------------- struct Int ---------------------
type Int int

func (a Int) Max(b Int) Int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func (i *Int) Set(a Int) {
	*i = a
}
func (i Int) Print() {
	fmt.Printf("value = %v\n", i)
}

// ---------------------- struct Map -------------------
type Map map[string]string
type iMap Map
type slice []int

func (m Map) Print() {
	for k, v := range m {
		fmt.Println(k + ":" + v)
	}
}
func (s slice) Print() {
	for k, v := range s {
		fmt.Printf("%d : %d\n", k, v)
	}
}

// ----------------------- struct T ---------------------
type T struct {
	a int
}

func (t *T) Set(i int) {
	t.a = i
}
func (t T) Get() int {
	return t.a
}
func (t *T) Print() {
	fmt.Printf("%p , %v, %d\n", t, t, t.a)
}
