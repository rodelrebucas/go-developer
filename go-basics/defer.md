A. Defer - executes a function before the main function returns
```
fmt.Println("start")
defer fmt.Println("middle")
fmt.Println("end")
// start
// end
// middle
```

B. Executes in LIFO
```
fmt.Println("start")
defer fmt.Println("middle")
fmt.Println("end")
// end
// middle
// start
```

C. Closing a resource before function returns
```
    import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
```

D. Defer process variables at the time it is called
```
a := "start"
fmt.Println(a)
a = "end"
// start
```
