package main

import "fmt"

func main() {
	// A. Creating empty interface
	var myObj writerCloser = myWriterCloser{}
	fmt.Println(myObj) // {}

}

// interfaces
type writer interface {
	Write([]byte) (int, error)
}

type closer interface {
	Close() error
}

type writerCloser interface {
	writer
	closer
}

// combined interfaces
type myWriterCloser struct{}

// implement the write method
func (wc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

// implement the close method
func (wc myWriterCloser) Close() error {
	return nil
}

// with close()
/** Output
I am tes
ting buf
fers
*/

// without close()
/** I am tes
ting buf
*/
