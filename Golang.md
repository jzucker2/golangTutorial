#Notes on Golang

## Starting points
- [ ] [A Tour of Go](https://tour.golang.org/welcome/1)
- [ ] [How to write Go Code](https://golang.org/doc/code.html)
- [ ] [Effective Go](https://golang.org/doc/effective_go.html)
- [ ] [Go's Declaration Syntax](https://blog.golang.org/gos-declaration-syntax)

## Cheat Sheet

* First [good cheat sheet](https://github.com/a8m/go-lang-cheat-sheet)

## Stray observations
* Need to understand this $GOPATH variable
* Exported names in packages begin with a capital letter, unexported names are lowercase (unexported means inaccessible/private)
* For functions, type comes after variable name
* print type: `fmt.Printf("x is of type %T\n", x)`
* Short assignment seems interesting and a possible minefield for me: `:=` with more notes [here](http://stackoverflow.com/questions/12611561/why-does-go-have-short-assignments-inside-functions)
* Read up more on [defer in this blog post](https://blog.golang.org/defer-panic-and-recover)
* Array size must be set with declaration and its size is constant and part of its type
* Slice literal is like an array literal without the length (don't have to specify length in the declaration)
* Slice has both _a length_ (number of elements in the slice) and _a capacity_ (number of elements in the underlying array counting from the first element in the slice)
  ```
  length := len(s)
  capacity := cap(s)
  ```
  
* What is `make`?
* [Slice capacity with odd numbers of elements differs from capacity with even numbers of elements](http://stackoverflow.com/questions/32995623/why-does-slice-capacity-with-odd-numbers-differ-from-behavior-with-even-numbers) and another explanation [here](http://stackoverflow.com/questions/38543825/appending-one-element-to-nil-slice-increases-capacity-by-two)
* You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).
* Pointer receivers vs. value receivers?
* Interface is a tuple of `(value, type)`

### Extra resources

Here is a great [list of Go learning resources](https://github.com/golang/go/wiki/Learn)
