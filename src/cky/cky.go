package main

import "fmt"

func main() {

	fmt.Println("Hello 世界")

	var s *Student

	s = new(Student)

	s.name = "codekissyoung"

	s.name = "ok"

	fmt.Printf("type : %T value: %v", *s, *s)

}

type Student struct {
	name string
	age  int8
}
