```
package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

        fmt.Println(len(slice))

        extra := []int{4, 5, 6}
	slice = append(slice, extra...)
	fmt.Println(slice)

	slice = append([]int{0}, slice...)
	fmt.Println(slice)
        fmt.Println(len(slice))

        slice[0] = 99
        fmt.Println(slice)

        slice[7] = 99
        fmt.Println(slice)
}
```
