package main

import (
	"bytes"
	"fmt"
)

func main() {
	// A. Creating interface with more than one methods
	var wc writerCloser = newBufferedWriterCloser() // Call the constsructor
	wc.Write([]byte("I am testing buffers"))
	wc.Close()
	// wc.buffer is not accessible
	// writerCLoser is not aware of the implementation of bufferedWriterCloser

	// type conversion with error handling to avoid panic
	bwc, ok := wc.(*bufferedWriterCloser)
	if ok {
		fmt.Println("Buffer: ", bwc.buffer) // buffer is now accessible
	} else {
		fmt.Println("Convertion failed")
	}

}

// interfaces
type writer interface {
	Write([]byte) (int, error)
}

type closer interface {
	Close() error
}

// combined interfaces
type writerCloser interface {
	writer
	closer
}

type bufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// implement the write method
func (bwc *bufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	// read and print only 8 characters
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// implement the close method
func (bwc *bufferedWriterCloser) Close() error {
	// read the next n bytes from the buffer
	// else return all bytes if n < buffer length
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// constructor
func newBufferedWriterCloser() *bufferedWriterCloser {
	return &bufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
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
