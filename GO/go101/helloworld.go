package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}
func pointer() string {
	test := 101
	fmt.Printf("Hello World %d \n", test)
	fmt.Printf("Result : %d\n", add(3, 4))
	i := 7
	j := 10
	var p *int
	p = &i
	fmt.Printf("address i : %v\n", &i)
	fmt.Printf("address j : %v\n", &j)
	fmt.Printf("P : %v \n", p)
	fmt.Printf("*P : %v \n", *p)
	*p = 6
	fmt.Printf("*p : %d\n", *p)
	fmt.Printf("i : %d\n", i)
	fmt.Printf("address p : %v\n", p)
	p = &j
	fmt.Printf("p : %v\n", p)
	fmt.Printf("*p : %d\n", *p)
	fmt.Printf("address p : %v\n", p)
	return "Pointer & Address Topic\n-----------------\n"
}

func rountine() string {
	go f1()
	go f2()
	fmt.Println("main_1")
	time.Sleep(time.Second * 2)
	fmt.Println("main_2")
	return "Rountine Topic\n-----------------\n"
}

func channel() string {
	// runtime.GOMAXPROCS(1)
	go pong()
	fmt.Println("main_1_1")
	c <- "ping"
	c <- "aaa"
	fmt.Println("main_1_2")
	return "Channel Topic\n-----------------\n"
}

func main() {
	pointer := pointer()
	fmt.Println(pointer)

	rountine := rountine()
	fmt.Println(rountine)

	channel := channel()
	fmt.Println(channel)
}

func f1() {
	fmt.Println("f1_1")
	time.Sleep(time.Second * 1)
	fmt.Println("f1_2")
}

func f2() {
	fmt.Println("f2")
}

var c = make(chan string)

func pong() {
	var str string
	for {
		str = <-c
		fmt.Println(str)
		if str == "ping" {
			fmt.Println("Pingg")
		} else {
			fmt.Println("err")
		}
	}
}
