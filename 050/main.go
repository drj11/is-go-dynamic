package main

import (
	"fmt"
)

func main() {
	set := map[int]struct{}{}
	fmt.Println(set)

	set[2] = struct{}{}
	set[3] = struct{}{}
	fmt.Println(set)

	delete(set, 3)
	fmt.Println(set)

        _, ok := set[3]
        if ok {
          fmt.Println("in")
        } else {
          fmt.Println("out")
        }
}
