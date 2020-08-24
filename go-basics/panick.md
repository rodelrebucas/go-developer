A. Panick - when application can not continue to function / execute - unrecoverable events

```
// import fmt and panic
fmt.Println("start")
panic("Something bad happened")
fmt.Println("end") // unreachable
// start
// panic: Something bad happened
// some detailed errors here...

```

B. `panic` executes after `defer`

C. Use `recover` to recover from `panic`. Use only inside `defer` functions

```
// import fmt, panic, log
func main() {
    fmt.Println("start")
    // higher function in call stack will continue unless panic is re throwed
    panicker()
    fmt.Prinln("end")
}

func panicker() {
    // this function will panic and not continue
    fmt.Println("panic start")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error: ", err) // gets the panic error
            // with re-panic
            // panic(err) - this will re throw the panic and stops the caller
            // start
            // panic start
            // Error: ...
            // Panic details ...
        }
    }()
    panic("something bad happened")
    fmt.Println("done panicking")
}
// without re panic
// start
// panic start
// Error: ...
// end
```
