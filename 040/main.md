```
// 040/main.go
package main

import (
	"fmt"
)

func main() {
	spell := map[int]string{1: "one", 2: "two", 4: "four"}
	fmt.Println(spell[1], spell[2], spell[3], spell[4])

	three, ok := spell[3]
	fmt.Println(three, ok)

	fmt.Println(len(spell))
}
```
