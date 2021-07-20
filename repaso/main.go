package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 9
	y := 0

	fmt.Println(x, y)
	myValue, err := strconv.ParseInt("20", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Array
	m := make(map[string]int)
	m["key"] = 90

	fmt.Println(m["key"])
	fmt.Println("================")

	// Slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index, value)
	}

	s = append(s, 76)
	for index, value := range s {
		fmt.Println(index, value)
	}

	// c := make(chan int)
	// go doSomething(c)
	// <-c

	fmt.Println("================")
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	c <- 1
}
