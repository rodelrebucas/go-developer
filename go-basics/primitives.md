A. Uninitialized boolean type variable contains false value

```
thisIsTrue := true
var n bool
fmt.Println(n) // false
```

B. Signed Integers

Types:

signed:

    int, int8, int16, int32, int64

unsigned:

    uint, uint8, uint16, unint32

byte - is an alias for 8bit unsigned integer so unint8 == byte

\*Operation on different types is invalid

C. Bit operators

    AND, OR, EXCLUSIVE OR, AND NOT

```
a := 10 // 1010
b := 3 //  0011

fmt.Println(a & b) // 2 = 0010
fmt.Println( a | b) // 11 = 1011
fmt.Println(a ^ b) // 9 = 1001
fmt.Println(a &^ b) // 8 = 1000
// AND NOT
// a AND NOT b
// 1010 & 1100 = 1000
```

D. Bit shifting

```
a := 8 // 2^3 = 1000
fmt.Println( a << 3) // 64 = 2^3 * 2^3 = 100 0000
fmt.Println( a >> 3) // 1 = 2^3 / 2^3 = 0001
```

E. Floating Numbers: `float32 and float64`

```
var b float32 = 3.14 // float32
v := 13.7e72 // float64
```

F. Complex types

```
var n complex64 = 1 + 2i // (1 + 2i)
```

F. Text

String - any utf8 character, read-only slice of bytes

```
s := "String example"
fmt.Println(s) // "String example"
fmt.Println(s[2]) // "114"
fmt.Printf("%T", s[2]) // uint8
fmt.Printf("%v", strin(s[s])) // r
```

Operations

1. Concatenations

```
s1 := "String"
s2 := "Example"
fmt.Printf("%v", s1 + s2) // StringExample
```

2. Convert to slice of bytes

```
s1 := "String"
s2 := []byte(s1)
fmt.Printf("%v", s2) // [83 116 114 105 110 103]
```

G. Rune - alias type of int32 , also represents a Unicode code point

```
 r := 'a' // 97
 var rn rune := 'a' // 97
```
