# Is Go a Dynamic Language?

## Abstract

Can Go replace your favourite dynamic language?
Go is a systems programming language developed by Google and in use
at ScraperWiki. Here I'll introduce Go in the context of "what
about Go is dynamic?". There will be some live coding using
play.golang.org. I'll also briefly outline what I've had
experience of making things using Go (at ScraperWiki).

## Brief Introduction to drj

Programming C, Python, /bin/sh, Lua, JavaScript, Go in some
combination for over 20 years.

First Engineer at ScraperWiki

ScraperWiki data rich web applications
Liverpool
Since 2009

## Brief Introduction to Go

- Systems programming language
- C like in syntax
- Static compilation
- All module resolution / linking done at compilation time
- Compiles to single binary

- Resizable datastructures: slice (aka list/array); map (aka dict/hash)
- First class functions
- "infinite" number of Goroutines (similar to greenthreads / coroutines)
- goroutines communicate via chan
- Automatic Garbage Collection

- Polymorphism via interface types
- dynamic type switch
- plenty of reflection; enough to implement JSON marshalling

## Civil Service People Survey

A major bespoke customer solution: a web app
used by thousands of users to query, display, and help analyse survey
results from a people survey in an organisation with 500,000 employees.
Frontend in Python, backend API in Go. About 7000 lines of code.

It has some novel features:

30 million row dataset (a few GB per year) is downloaded from S3 and
parsed into an internal indexed datastructure every time the
server starts. Using multiple CPUs to do this in parallel means
that all 7 years can be loaded in about 50 seconds.

Queries are computed every time and consider "every row".
Response times are in the few millisecond range for the worst
queries (down from 10 minutes with our protoype in Python /
PostgreSQL). Probably would have had to cache / precompute if
doing it in Python (which an earlier version did).

Go compiles to a static binary and the Docker image that runs
the API has only 4 files in its filesystem: our compiled binary,
/etc/passwd and /etc/group, and a ca-certificates file (and
actually 2 more that I just discovered: a redundant .gitignore
and a Dockerfile). The data processed by the API is quite sensitive,
IL2, so we have to take care in safeguarding the data. We regard Go's
ability to compile down to very nearly standalone binaries
as one of the components in our security fence: even if you
manage to escape the binary, you're inside a docker container
with only one binary file and no system utilities.

## pdftables.com

pdftables.com is a relatively new product from ScraperWiki; the
frontend is written in Go, with other parts written in Go and
Python.


## Hanoverd

https://github.com/scraperwiki/hanoverd

Allows seamless switching from one deployed version of a
dockerised application to another.

And other devops tools.

Many little devops tools, some of which might have wider
application.

A few devops server components that manage and orchestrate our
seamless build and deployment systems.

Some code is available on github.com

## Hello World


## Resizable arrays [030]

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

(notice in this example, the difference between initialisation
and assignment)

You can prepend items to a slice, using... append()!:

```
slice = append([]int{0}, slice...)
```

(note that this is a one element slice that is being prepended,
any length slice can be used)

A slice is indexed using square brackets, to get or set the
element:

```
slice[0] = 99
```

Slices are 0-based. Length is obtained with the len() function.

Using an inappropriate index is a runtime check and will result in a
panic.  Generally this quits your program and dumps all stacks. But if
you're writing a webserver, the server loop recovers from the
panic and continues serving, so that the panic is isolated to
the web request that caused it.

## Resizable hash tables

The hash-table-like datastructure is called the map, as in
associative map. It's similar to dict() from Python:

```
spell := map[int]string{1: "one", 2: "two", 4: "four"}
fmt.Println(spell[1], spell[2], spell[3], spell[4])
```

Index using square brackets, like slices. Unlike slices, trying
to access a map entry that does not exist does not cause a
runtime panic. Instead we get an empty string. Well we get a
string in this case because the map is typed to return string.
The empty string is the *zero value* for the string type. Every
type has a *zero value*, which is 0 for numeric types, false
for the bool type, the nil map for all maps, and so on.

How do we tell the difference between having a map entry that is
an empty string, and having no map entry at all? When assigning
the result of a map lookup, Go has a special form with an extra
variabe (traditionally called ok):

```
three, ok := spell[3]
fmt.Println(three, ok)
```

Like slices, the number of elements in the map is returned by the len
function:

```
fmt.Println(len(spell))
```


## sets



## defer()
