A. Slice - Describes a part of an _array_. Copied by reference

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

B. Append - append(parentSlice, elements)

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
