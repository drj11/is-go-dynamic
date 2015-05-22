# Is Go a Dynamic Language?

## Abstract

I will briefly introduce Go, a systems programming language
developed by Google and in use at ScraperWiki. Outline some of
the ways in which Go seems like a dynamic language, and some of
the ways in which Go seems like a static language. The
presentation will be largely illustrated with coded examples.

## Random idea for presenting

- Use shortcode links to play.golang.org
- put the links in the slides
- Just make the entire presentation a series of tabs in the browser

## Brief Introduction to Go

- Systems programming language
- Static compilation
- All module resolution / linking done at compilation time
- Produces single binary (nearly static)

- Resizable datastructures: slice (aka list/array); map (aka dict/hash)
- "infinite" number of Goroutines (similar to greenthreads / coroutines)
- First class functions
- Automatic Garbage Collection

- Polymorphism via interface types
- dynamic type switch

- plenty of reflection; enough to implement JSON marshalling

## What are ScraperWiki using it for?

One major bespoke customer solution: a web app used by thousands
of users to query, display, and help analyse survey results from
a people survey in an organisation with 500,000 employees.

Many little devops tools, some of which might have wider
application.

A few devops server components that manage and orchestrate our
seamless build and deployment systems.

Some code is available on github.com

## Type Declarations

No deed to declare types:


```
package main

import (
        "fmt"
        "os"
)

func main() {
        a := os.Args[1]
        fmt.Println(a)
}
```

But this is not because variables can hold variables of any type, but
because the type of the variable can often be _inferred_. In this case
`a` is inferred to be of type `string`. Subsequently trying to
store a non-string value will fail:


```
        a = 4
```

But notice that the failure is at compile time, not runtime.

## Resizable arrays

The array-like datastructure is called the _slice_. There is
some scope for confusion here, as Go also has a datastructure
called an _array_. The Go array is not resizable, it is a linear
sequence of elements, all of the same type, and the length of
which is fixed at compile time. Every Go slice is backed by a Go
array. You can use Go arrays directly, but most of the time you
use a slice.

You can append items to a slice, using append():

```
slice := []int{1, 2, 3}
slice = append(slice, 4, 5, 6)
```

You can extend a slice with another slice, using append():

```
extra := []int{4,5,6}
slice = append(slice, extra...)
```

You can prepend items to a slice, using... append()!:

```
slice = append([]int{0}, slice...)
```

## Resizable hash tables

The hash-table-like datastructure is called the map, as in
associative map. It's similar the dict() from Python:

```
spell := string[int]{1:"one", 2:"two", 3:"three"}
fmt.Println(spell[1], spell[2], spell[7])
```


