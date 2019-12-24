package main
import "fmt"
type Vertex struct {
	Lat, Long float64
}

// 返回局部变量
func f() *int {
	v := 1
	return &v
}

// 指针入参
func incr(p *int) int {
	*p++
	return *p
}

// new 语法糖 返回指针
func newInt() *int {
	return new(int)
}

func main() {

	a:=[3]int{}

	fmt.Println(a)

	v := 1

	incr( &v )
	incr( &v )

	o := newInt()

	*o = 100

	fmt.Println( o )
	fmt.Println( v )
}
