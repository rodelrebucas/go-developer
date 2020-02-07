A. Interface - collection of function signatures; describes _behavior_ unlike _struct_ which which describes data

B. Naming convention should be _method name_ + _er_

```
func main() {
	// create variable of type Writer
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	// describe the method signature
	Write([]byte) (int, error)
}


```

C. Implicitly implements the Writer interface by creating the method signature _Write_

```
type ConsoleWriter struct{}

func (vw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
```

D. Using custom type other than struct

```
func main() {
	myInt := myInt(10)
	var doubled incrementer = &myInt
	fmt.Printf("%v\n", myInt)            // 10
	fmt.Printf("%v\n", doubled.Double()) // 20
	fmt.Printf("%v", doubled.Double())   // 40

}

type incrementer interface {
	Double() int
}

type myInt int

func (ic *myInt) Double() int {
	*ic = *ic * 2
	return int(*ic)
}
```
