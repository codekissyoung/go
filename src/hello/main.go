package main

import (
    "errors"
    "fmt"
    "time"
)

// ---------------------------------- struct -----------------------------------

type user struct {
    name string
    age  byte
}

type manager struct {
    user
    title string
}

type MyInt int

type Printer interface {
    Print()
}

// ---------------------------------- main ------------------------------------

func main() {

    // test_slice()

    // test_map()

    // test_struct()

    // test_interface()

    // test_method()


    /*
    go task(1)

    go task(2)

    time.Sleep(time.Second * 6)
    */

    done := make(chan bool)
    data := make(chan int)

    go consumer(data, done)
    go producer(data)

    <-done
}

// ------------------------------------ func -------------------------------------

// test multi return
func div( a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// return func type
func test(x int) func() {
    return func() {
        println(x)
    }
}

// slice
func test_slice() {

    x := make( []int, 0, 5 )

    for i:= 0; i < 8; i++ {
        x = append(x, i)
    }

    fmt.Println(x)
}

// map
func test_map() {

    m := make(map[string]int)

    m["a"] = 1

    fmt.Println(m)

    x, ok := m["b"]

    fmt.Println(x, ok)

    delete(m, "a")

}

// struct
func test_struct() {
    var m manager

    m.name  = "Tom"
    m.age   = 29
    m.title = "CTO"

    // fmt.Println(m)
    println(m.ToString())

}

// method
func (x *MyInt) inc() {
    *x++
}

func test_method() {

    var x MyInt = 10

    x.inc()

    println(x)

}

func (u user) ToString() string {

    return fmt.Sprintf("%+v", u)
}

func (u user) Print() {
    fmt.Printf("%+v\n", u)
}

func test_interface () {
    var u user;
    u.name = "Alice"
    u.age  = 30

    var p Printer = u

    p.Print()
}

func task(id int) {
    for i :=0; i < 5; i++ {
        fmt.Printf("%d : %d\n", id, i)
        time.Sleep(time.Second)
    }
}


func consumer( data chan int, done chan bool ) {

    for x := range data {
        println("recv:", x)
    }

    done <- true
}

func producer( data chan int ) {
    for i := 0; i < 4; i++ {
        data <- i
    }
    close(data)
}


