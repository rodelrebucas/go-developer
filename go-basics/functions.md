A. Passing multiple arguments with the same type
```
greet("Hi", "Ed") // Hi Ed

func greet(msg, name string) { // parameter, parameter type
	fmt.Println(msg, name)
}
```

B. Passing pointer type argument
```
name := "Edy"
msg := "Hello"
greet2(&msg, &name) // Hello Edy

func greet2(msg, name *string) { // parameter, parameter type
	fmt.Println(*msg, *name)
}
```

C. Variatic parameter
```
greet3("Hello", "Ed", "Edy") // Hello [Ed Edy]

func greet3(msg string, names ...string) {
	fmt.Println(msg, names)
}

```

D. Return value
```
fmt.Println(greet4("Hellooo")) // Echoed Hellooo


func greet4(msg string) string {
	return "Echoed " + msg
}
```

E. Pointer return value is created on Heap rather than the stack,
since variables in the function is created on that function's stack
and deleted after the fn finishes execution  or  returns

E. Named return value
```
fmt.Println(greet5("Heyy")) // Heyy echoed

func greet5(msg string) (echo string) {
	echo = msg + " " + "echoed"
	return
}
```

F. Returning value and error