A. specifc size
```
// [size]type
grades := [3]int{1, 2, 3}
fmt.Println(grades) // [1 2 3]
```

B. Without specifying the size
```
// [...]type
grades2 := [...]int{1, 2, 3, 4}
fmt.Println(grades2) // [1 2 3 4]
```

C. Empty array and assigning value later
```
var students [3]string
fmt.Println(students) // []
students[0] = "student1"
fmt.Println(students) // [student1 ]
fmt.Println(len(students)) // 3
```

D. Arrays are copied by value
```
a := [...]int{1, 2, 3}
b := a
b[1] = 200
fmt.Println(a) // [1 2 3]
fmt.Println(b) // [1 200 3]
```

E. Using pointers to access a reference to array's memory address
```
a := [...]int{1, 2, 3}
b := &a
b[1] = 200
fmt.Println(a) // [1 200 3]
fmt.Println(*b) // [1 200 3]
```

F. Slice - Copied by reference
```
a := []int{1, 2, 3, 5, 6}
b := a
b[1] = 5
c := a[2:]
fmt.Println(a)      // [1 5 3]
fmt.Println(len(a)) // 3
fmt.Println(cap(a)) // 3
fmt.Println(c)      // [3 5 6]
```

G. Make - creates array with zeroed values
```
// make(type, size, capacity)
a := make([]int, 3, 100)
fmt.Println(a)             // [0 0 0]
fmt.Printf("%v\n", len(a)) // 3
fmt.Printf("%v\n", cap(a)) // 100
```

H. Append 
```
a := []int{}
fmt.Println(a) // []
fmt.Println(len(a)) // 0
fmt.Println(cap(a)) // 0
a = append(a, 1, 2)
fmt.Println(a) // [1 2]
fmt.Println(len(a)) // 2
fmt.Println(cap(a)) // 2
```