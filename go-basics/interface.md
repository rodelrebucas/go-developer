

A. Interface - collection of function signatures; describes *behavior* unlike *struct* which which describes data

B. Naming convention should suffix with *er* with single method

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
C. Implicitly implements the Writer interface by  creating the method signature *Write*

```
type ConsoleWriter struct{}

func (vw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
```
