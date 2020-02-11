A. map[< key typed >]< value typed >{}

```
countryPopulations := map[string]int{
    "countryA": 100,
    "countryB": 200,
    "countryC": 300,
}
fmt.Println(countryPopulations) // map[countryA:100 countryB:200 countryC:300]
```

B. instantiating the map without values

```
options := make(map[string]string)
options = map[string]string{
    "a": "option A",
    "b": "option B",
}
fmt.Println(options)
```

C. Using maps

```
fmt.Println(options["a"]) // Option A
```

D. Adding other key

```
options["c"] = "option C"
fmt.Println(options["c"]) // Option C
```

E. Deleting value

```
delete(options, "a")
```

F. Checking if value exist

```
value, ok := options["b"]
fmt.Println(value, ok) // option B true
value, ok = options["a"]
fmt.Println(value, ok) //   false
```

G. Maps are copied by reference

```
options2 := options
fmt.Println(options2) // map[b: option c:option C]
```
